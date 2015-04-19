package kafka

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PieterD/kafka-processor/killchan"
	"github.com/Shopify/sarama"
	"github.com/samuel/go-zookeeper/zk"
)

var ErrNoBrokers = errors.New("No kafka brokers found")
var ErrAlreadyListening = errors.New("Already listening to stream")
var ErrNotListening = errors.New("Not listening to stream")
var ErrListenerQuit = errors.New("Listener exited unexpectedly; offset out of range?")

type Message struct {
	Key       []byte
	Val       []byte
	Topic     string
	Partition int32
	Offset    int64
}

type Stream struct {
	Topic     string
	Partition int32
}

type Kafka struct {
	isclosed    int32
	logger      *log.Logger
	zkpeers     []string
	zooConn     *zk.Conn
	kfkConn     sarama.Client
	kfkConsumer sarama.Consumer
	kfkProducer sarama.SyncProducer
	lock        sync.Mutex
	listenbus   chan listenRequest
	messagebus  chan Message
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

func (k *Kafka) run() {
	defer k.close()
	consumers := make(map[Stream]listener)
	for {
		select {
		case req := <-k.listenbus:
			if req.response == nil {
				err := k.exited(consumers, req.stream)
				if err != nil {
					//TODO: This needs to be better.
					panic(err)
				}
				continue
			}
			if req.enable {
				err := k.listen(consumers, req.stream, req.offset)
				if err != nil {
					req.response <- err
				}
			} else {
				err := k.unlisten(consumers, req.stream)
				if err != nil {
					req.response <- err
				}
			}
			close(req.response)
		}
	}
}

func (k *Kafka) exited(consumers map[Stream]listener, stream Stream) error {
	_, ok := consumers[stream]
	if ok {
		// Listener quit without us first removing it from the map;
		// this means it was not caused by us, and thus unexpected.
		return ErrListenerQuit
	}
	return nil
}

func (k *Kafka) listen(consumers map[Stream]listener, stream Stream, offset int64) error {
	_, ok := consumers[stream]
	if ok {
		return ErrAlreadyListening
	} else {
		conn, err := k.kfkConsumer.ConsumePartition(stream.Topic, stream.Partition, offset)
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
			defer func() { k.listenbus <- listenRequest{stream: stream} }()
			defer l.killed.Kill()
			for msg := range conn.Messages() {
				k.messagebus <- Message{
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

func (k *Kafka) unlisten(consumers map[Stream]listener, stream Stream) error {
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

func New(logger *log.Logger, zkpeers []string) (*Kafka, error) {
	k := new(Kafka)
	k.listenbus = make(chan listenRequest)
	k.messagebus = make(chan Message)
	k.logger = logger
	k.zkpeers = zkpeers
	err := k.connect()
	if err != nil {
		return nil, err
	}
	go k.run()
	return k, nil
}

func (k *Kafka) connect() error {
	zooConn, _, err := zk.Connect(k.zkpeers, time.Second)
	if err != nil {
		return fmt.Errorf("Failed to connect to zookeeper: %v", err)
	}
	k.zooConn = zooConn

	kafkaBrokers, err := k.GetKafkaBrokersFromZookeeper()
	if err != nil {
		k.close()
		return fmt.Errorf("Failed to feth brokers from zookeeper: %v", err)
	}
	if len(kafkaBrokers) == 0 {
		k.close()
		return ErrNoBrokers
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "kafka-go"

	kfkConn, err := sarama.NewClient(brokerStrings(kafkaBrokers), config)
	if err != nil {
		k.close()
		return fmt.Errorf("Failed to connect to Kafka: %v", err)
	}
	k.kfkConn = kfkConn

	consumer, err := sarama.NewConsumerFromClient(kfkConn)
	if err != nil {
		k.close()
		return fmt.Errorf("Failed to create Kafka consumer: %v", err)
	}
	k.kfkConsumer = consumer

	producer, err := sarama.NewSyncProducerFromClient(kfkConn)
	if err != nil {
		k.close()
		return fmt.Errorf("Failed to create sarama syncproducer: %v", err)
	}
	k.kfkProducer = producer

	return nil
}

func (k *Kafka) close() {
	if atomic.CompareAndSwapInt32(&k.isclosed, 0, 1) {
		if k.kfkProducer != nil {
			k.kfkProducer.Close()
		}
		if k.kfkConsumer != nil {
			k.kfkConsumer.Close()
		}
		if k.kfkConn != nil {
			k.kfkConn.Close()
		}
		if k.zooConn != nil {
			k.zooConn.Close()
		}
	}
}

func (k *Kafka) Listen(topic string, partition int32, offset int64) error {
	stream := Stream{Topic: topic, Partition: partition}
	req := listenRequest{
		stream:   stream,
		offset:   offset,
		response: make(chan error),
		enable:   true,
	}
	k.listenbus <- req
	err, ok := <-req.response
	if ok {
		return err
	}

	return nil
}

func (k *Kafka) Unlisten(topic string, partition int32) error {
	stream := Stream{Topic: topic, Partition: partition}
	req := listenRequest{
		stream:   stream,
		response: make(chan error),
		enable:   false,
	}
	k.listenbus <- req
	err, ok := <-req.response
	if ok {
		return err
	}

	return nil
}
