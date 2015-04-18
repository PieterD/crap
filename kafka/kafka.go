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
	consumers   map[Stream]*partConsumer
	messagebus  chan Message
}

func New(logger *log.Logger, zkpeers []string) (*Kafka, error) {
	k := new(Kafka)
	k.consumers = make(map[Stream]*partConsumer)
	k.messagebus = make(chan Message)
	k.logger = logger
	k.zkpeers = zkpeers
	err := k.connect()
	if err != nil {
		return nil, err
	}
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
		k.Close()
		return fmt.Errorf("Failed to feth brokers from zookeeper: %v", err)
	}
	if len(kafkaBrokers) == 0 {
		k.Close()
		return ErrNoBrokers
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "kafka-go"

	kfkConn, err := sarama.NewClient(brokerStrings(kafkaBrokers), config)
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to connect to Kafka: %v", err)
	}
	k.kfkConn = kfkConn

	consumer, err := sarama.NewConsumerFromClient(kfkConn)
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to create Kafka consumer: %v", err)
	}
	k.kfkConsumer = consumer

	producer, err := sarama.NewSyncProducerFromClient(kfkConn)
	if err != nil {
		consumer.Close()
		return fmt.Errorf("Failed to create sarama syncproducer: %v", err)
	}
	k.kfkProducer = producer

	return nil
}

func (k *Kafka) Close() error {
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
	return nil
}

type partConsumer struct {
	conn   sarama.PartitionConsumer
	offset int64
	stream Stream
	kill   killchan.Killchan
	killed killchan.Killchan
}

func (k *Kafka) Listen(topic string, partition int32, offset int64) error {
	stream := Stream{Topic: topic, Partition: partition}

	// Check if the stream was already added
	k.lock.Lock()
	_, ok := k.consumers[stream]
	k.lock.Unlock()
	if ok {
		return ErrAlreadyListening
	}

	// Create parititon consumer
	kfkPartConsumer, err := k.kfkConsumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		return err
	}

	pc := &partConsumer{
		conn:   kfkPartConsumer,
		offset: offset,
		stream: stream,
		kill:   killchan.New(),
		killed: killchan.New(),
	}

	// Re-check if stream was added in the mean time
	k.lock.Lock()
	_, ok = k.consumers[stream]
	if !ok {
		k.consumers[stream] = pc
	}
	k.lock.Unlock()

	// It was added in the mean time; close the partition consumer
	if ok {
		err = kfkPartConsumer.Close()
		if err != nil {
			//TODO: Is this the right thing to do?
			return err
		}
		return ErrAlreadyListening
	}

	go pc.run(k)

	return nil
}

func (k *Kafka) Unlisten(topic string, partition int32) error {
	k.lock.Lock()
	pc, ok := k.consumers[stream]
	k.lock.Unlock()
	if !ok {
		return ErrNotListening
	}

	pc.close()

	k.lock.Lock()
	delete(k.consumers, stream)
	k.lock.Unlock()

	return nil
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
