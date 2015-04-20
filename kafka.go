package kafka

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"

	"github.com/PieterD/kafka/internal/listenhandler"
	"github.com/PieterD/kafka/internal/transmithandler"
	"github.com/PieterD/kafka/internal/zoohandler"
	"github.com/Shopify/sarama"
)

var ErrNoBrokers = errors.New("No kafka brokers found")
var ErrTransmitHandlerClosed = errors.New("Transmit handler has closed")

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
	zh       *zoohandler.ZooHandler
	kfkConn  sarama.Client
	lh       *listenhandler.ListenHandler
	incoming chan Message
	th       *transmithandler.TransmitHandler
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
	k.incoming = make(chan Message)
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
	zh, err := zoohandler.New(k.zkpeers, k.logger)
	if err != nil {
		return fmt.Errorf("Failed to connect to zookeeper: %v", err)
	}
	k.zh = zh

	kafkaBrokers, err := k.zh.GetKafkaBrokers()
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to fetch brokers from zookeeper: %v", err)
	}
	if len(kafkaBrokers) == 0 {
		k.Close()
		return ErrNoBrokers
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = clientid

	kfkConn, err := sarama.NewClient(kafkaBrokers.Strings(), config)
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

	k.th, err = transmithandler.New(kfkConn)
	if err != nil {
		k.Close()
		return fmt.Errorf("Failed to start transmit handler: %v", err)
	}

	return nil
}

func (k *Kafka) Close() {
	if atomic.CompareAndSwapInt32(&k.isclosed, 0, 1) {
		if k.th != nil {
			k.th.Close()
		}
		if k.lh != nil {
			k.lh.Close()
		}
		if k.kfkConn != nil {
			k.kfkConn.Close()
		}
		if k.zh != nil {
			k.zh.Close()
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

func (k *Kafka) Outgoing() <-chan transmithandler.Transmit {
	return k.th.Outgoing()
}

func (k *Kafka) Send(key, val []byte, topic string) (partition int32, offset int64, err error) {
	trans, ok := <-k.th.Outgoing()
	if !ok {
		return 0, 0, ErrTransmitHandlerClosed
	}
	return trans.Send(key, val, topic)
}
