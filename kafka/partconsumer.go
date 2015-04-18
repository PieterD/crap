package kafka

import (
	"github.com/PieterD/kafka-processor/killchan"
	"github.com/Shopify/sarama"
)

type partConsumer struct {
	conn   sarama.PartitionConsumer
	stream Stream
	offset int64
	kill   killchan.Killchan
	killed killchan.Killchan
}

func newPartConsumer(consumer sarama.PartitionConsumer, stream Stream, offset int64) *partConsumer {
	return &partConsumer{
		conn:   consumer,
		stream: stream,
		offset: offset,
		kill:   killchan.New(),
		killed: killchan.New(),
	}
}

func (pc *partConsumer) close() bool {
	if pc.kill.Kill() {
		pc.killed.Wait()
		return true
	}
	return false
}

func (pc *partConsumer) run(k *Kafka) {
	defer pc.killed.Kill()
	defer pc.conn.Close()
	for {
		var msg Message

		// Receive a message
		select {
		case sMessage := <-pc.conn.Messages():
			msg = Message{
				Key:       sMessage.Key,
				Val:       sMessage.Value,
				Topic:     pc.stream.Topic,
				Partition: pc.stream.Partition,
				Offset:    sMessage.Offset,
			}
		case <-pc.kill.Chan():
			return
		}

		// Send the message on
		select {
		case k.messagebus <- msg:
		case <-pc.kill.Chan():
			return
		}
	}
}
