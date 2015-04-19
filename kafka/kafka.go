package kafka

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/PieterD/kafka-processor/kafka/internal/listenhandler"
	"github.com/Shopify/sarama"
	"github.com/samuel/go-zookeeper/zk"
)

var ErrNoBrokers = errors.New("No kafka brokers found")

type Message struct {
	Key       []byte
	Val       []byte
	Topic     string
	Partition int32
	Offset    int64
}

type Kafka struct {
	isclosed int32
	logger   *log.Logger
	zkpeers  []string
	zooConn  *zk.Conn
	kfkConn  sarama.Client
	lh       *listenhandler.ListenHandler
	incoming chan Message
}

type messageTransmitter chan<- Message

func (mt messageTransmitter) Transmit(key, val []byte, topic string, partition int32, offset int64) {
	mt <- Message{
		Key:       key,
		Val:       val,
		Topic:     topic,
		Partition: partition,
		Offset:    offset,
	}
}

func (mt messageTransmitter) Close() {
	close(mt)
}

func New(clientid string, logger *log.Logger, zkpeers []string) (*Kafka, error) {
	k := new(Kafka)
	k.logger = logger
	k.zkpeers = zkpeers
	err := k.connect(clientid)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (k *Kafka) Incoming() <-chan Message {
	return k.incoming
}

func (k *Kafka) connect(clientid string) error {
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
	config.ClientID = clientid

	kfkConn, err := sarama.NewClient(brokerStrings(kafkaBrokers), config)
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to connect to Kafka: %v", err)
	}
	k.kfkConn = kfkConn

	k.lh, err = listenhandler.New(kfkConn, messageTransmitter(k.incoming))
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to start listen handler: %v", err)
	}

	return nil
}

func (k *Kafka) Close() {
	if atomic.CompareAndSwapInt32(&k.isclosed, 0, 1) {
		if k.lh != nil {
			k.lh.Close()
		}
		if k.kfkConn != nil {
			k.kfkConn.Close()
		}
		if k.zooConn != nil {
			k.zooConn.Close()
		}
	}
}

func (k *Kafka) ListenNewest(topic string, partition int32) error {
	return k.lh.Listen(topic, partition, sarama.OffsetNewest)
}

func (k *Kafka) ListenOldest(topic string, partition int32) error {
	return k.lh.Listen(topic, partition, sarama.OffsetOldest)
}

func (k *Kafka) Listen(topic string, partition int32, offset int64) error {
	return k.lh.Listen(topic, partition, offset)
}

func (k *Kafka) Unlisten(topic string, partition int32) error {
	return k.lh.Unlisten(topic, partition)
}
