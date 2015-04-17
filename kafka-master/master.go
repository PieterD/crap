package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/samuel/go-zookeeper/zk"
)

var (
	fPeers = flag.String("peers", os.Getenv("ZOOKEEPER_PEERS"), "List of Zookeeper peer addresses (Defaults to ZOOKEEPER_PEERS env)")

	logger = log.New(os.Stderr, "master", log.LstdFlags)
)

func flagbad(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *fPeers == "" {
		flagbad("No zookeeper peers provided; use -peers or the ZOOKEEPER_PEERS environment variable")
	}

	zooConn, _, err := zk.Connect(strings.Split(*fPeers, ","), time.Second)
	if err != nil {
		logger.Panicf("Failed to connect to zookeeper: %v", err)
	}
	defer zooConn.Close()

	kafkaBrokers := GetKafkaBrokers(zooConn)
	if len(kafkaBrokers) == 0 {
		logger.Panicf("No kafka brokers found!")
	}
	for _, broker := range kafkaBrokers {
		fmt.Printf("broker %3d, addr: %s\n", broker.Id, broker.Addr())
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "kafkaproc.master"
	client, err := sarama.NewClient(KafkaBrokerStrings(kafkaBrokers), config)
	if err != nil {
		logger.Panicf("Creating sarama client: %v", err)
	}
	defer client.Close()
}
