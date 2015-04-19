package kafka

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/Shopify/sarama"
	"github.com/samuel/go-zookeeper/zk"
)

var ErrNoBrokers = errors.New("No kafka brokers found")
var ErrAlreadyListening = errors.New("Already listening to stream")
var ErrNotListening = errors.New("Not listening to stream")
var ErrListenerQuit = errors.New("Listener exited unexpectedly; offset out of range?")
var ErrListenHandlerClosed = errors.New("Listen handler had been closed")

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
	isclosed int32
	logger   *log.Logger
	zkpeers  []string
	zooConn  *zk.Conn
	kfkConn  sarama.Client
	//kfkProducer sarama.SyncProducer
	lh *listenHandler
}

func New(logger *log.Logger, zkpeers []string) (*Kafka, error) {
	k := new(Kafka)
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

	//producer, err := sarama.NewSyncProducerFromClient(kfkConn)
	//if err != nil {
	//	k.close()
	//	return fmt.Errorf("Failed to create sarama syncproducer: %v", err)
	//}
	//k.kfkProducer = producer

	return nil
}

func (k *Kafka) close() {
	if atomic.CompareAndSwapInt32(&k.isclosed, 0, 1) {
		if k.lh != nil {
			k.lh.close()
		}
		//if k.kfkProducer != nil {
		//	k.kfkProducer.Close()
		//}
		if k.kfkConn != nil {
			k.kfkConn.Close()
		}
		if k.zooConn != nil {
			k.zooConn.Close()
		}
	}
}
