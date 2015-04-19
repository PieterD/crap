package kafka

import (
	"fmt"

	"github.com/PieterD/kafka-processor/killchan"
	"github.com/Shopify/sarama"
)

type listenHandler struct {
	listenbus  chan listenRequest
	messagebus chan Message
	consumer   sarama.Consumer
	kill       killchan.Killchan
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

func newListenHandler(kfkConn sarama.Client) (*listenHandler, error) {
	consumer, err := sarama.NewConsumerFromClient(kfkConn)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kafka consumer: %v", err)
	}
	lh := &listenHandler{
		listenbus:  make(chan listenRequest),
		messagebus: make(chan Message),
		consumer:   consumer,
	}
	go lh.run()
	return lh, nil
}

func (lh *listenHandler) close() {
	lh.kill.Kill()
}

func (lh *listenHandler) run() {
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

func (lh *listenHandler) exited(consumers map[Stream]listener, stream Stream) error {
	_, ok := consumers[stream]
	if ok {
		// Listener quit without us first removing it from the map;
		// this means it was not caused by us, and thus unexpected.
		return ErrListenerQuit
	}
	return nil
}

func (lh *listenHandler) listen(consumers map[Stream]listener, stream Stream, offset int64) error {
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
				lh.messagebus <- Message{
					Key:       msg.Key,
					Val:       msg.Value,
					Topic:     stream.Topic,
					Partition: stream.Partition,
					Offset:    offset,
				}
			}
		}()
	}
	return nil
}

func (lh *listenHandler) unlisten(consumers map[Stream]listener, stream Stream) error {
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

func (lh *listenHandler) Listen(topic string, partition int32, offset int64) error {
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

func (lh *listenHandler) Unlisten(topic string, partition int32) error {
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
