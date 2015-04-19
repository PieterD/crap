package listenhandler

import (
	"errors"
	"fmt"

	"github.com/PieterD/kafka-processor/killchan"
	"github.com/Shopify/sarama"
)

var ErrAlreadyListening = errors.New("Already listening to stream")
var ErrNotListening = errors.New("Not listening to stream")
var ErrListenerQuit = errors.New("Listener exited unexpectedly; offset out of range?")
var ErrListenHandlerClosed = errors.New("Listen handler had been closed")

type Transmitter interface {
	Transmit(key, val []byte, topic string, partition int32, offset int64)
	Close()
}

type Stream struct {
	Topic     string
	Partition int32
}

type ListenHandler struct {
	listenbus   chan listenRequest
	transmitter Transmitter
	consumer    sarama.Consumer
	kill        killchan.Killchan
	dead        killchan.Killchan
}

type listenRequest struct {
	stream   Stream
	offset   int64
	response chan error
	enable   bool
}

type listener struct {
	stream Stream
	conn   sarama.PartitionConsumer
	killed killchan.Killchan
}

func New(kfkConn sarama.Client, transmitter Transmitter) (*ListenHandler, error) {
	consumer, err := sarama.NewConsumerFromClient(kfkConn)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kafka consumer: %v", err)
	}
	lh := &ListenHandler{
		listenbus:   make(chan listenRequest),
		transmitter: transmitter,
		consumer:    consumer,
		kill:        killchan.New(),
	}
	go lh.run()
	return lh, nil
}

func (lh *ListenHandler) Close() {
	lh.kill.Kill()
	lh.dead.Wait()
}

func (lh *ListenHandler) run() {
	defer lh.transmitter.Close()
	defer lh.dead.Kill()
	defer lh.consumer.Close()
	consumers := make(map[Stream]listener)
	defer func() {
		for stream := range consumers {
			lh.unlisten(consumers, stream)
		}
		for req := range lh.listenbus {
			if req.response != nil {
				req.response <- ErrListenHandlerClosed
				close(req.response)
			}
		}
	}()

	for {
		select {
		case <-lh.kill.Chan():
			return
		case req := <-lh.listenbus:
			if req.response == nil {
				err := lh.exited(consumers, req.stream)
				if err != nil {
					//TODO: This needs to be better.
					panic(err)
				}
				continue
			}
			if req.enable {
				err := lh.listen(consumers, req.stream, req.offset)
				if err != nil {
					req.response <- err
				}
			} else {
				err := lh.unlisten(consumers, req.stream)
				if err != nil {
					req.response <- err
				}
			}
			close(req.response)
		}
	}
}

func (lh *ListenHandler) exited(consumers map[Stream]listener, stream Stream) error {
	_, ok := consumers[stream]
	if ok {
		// Listener quit without us first removing it from the map;
		// this means it was not caused by us, and thus unexpected.
		return ErrListenerQuit
	}
	return nil
}

func (lh *ListenHandler) listen(consumers map[Stream]listener, stream Stream, offset int64) error {
	_, ok := consumers[stream]
	if ok {
		return ErrAlreadyListening
	} else {
		conn, err := lh.consumer.ConsumePartition(stream.Topic, stream.Partition, offset)
		if err != nil {
			return err
		}
		l := listener{
			stream: stream,
			conn:   conn,
			killed: killchan.New(),
		}
		consumers[stream] = l
		go func() {
			defer func() { lh.listenbus <- listenRequest{stream: stream} }()
			defer l.killed.Kill()
			for msg := range conn.Messages() {
				lh.transmitter.Transmit(msg.Key, msg.Value, stream.Topic, stream.Partition, offset)
			}
		}()
	}
	return nil
}

func (lh *ListenHandler) unlisten(consumers map[Stream]listener, stream Stream) error {
	l, ok := consumers[stream]
	if !ok {
		return ErrNotListening
	} else {
		delete(consumers, stream)
		//TODO: Perhaps use asynch close?
		l.conn.Close()
		l.killed.Wait()
	}
	return nil
}

func (lh *ListenHandler) Listen(topic string, partition int32, offset int64) error {
	stream := Stream{Topic: topic, Partition: partition}
	req := listenRequest{
		stream:   stream,
		offset:   offset,
		response: make(chan error),
		enable:   true,
	}
	lh.listenbus <- req
	err, ok := <-req.response
	if ok {
		return err
	}

	return nil
}

func (lh *ListenHandler) Unlisten(topic string, partition int32) error {
	stream := Stream{Topic: topic, Partition: partition}
	req := listenRequest{
		stream:   stream,
		response: make(chan error),
		enable:   false,
	}
	lh.listenbus <- req
	err, ok := <-req.response
	if ok {
		return err
	}

	return nil
}
