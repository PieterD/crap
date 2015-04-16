package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	fPeers     = flag.String("peers", os.Getenv("KAFKA_PEERS"), "List of Kafka peer addresses (Defaults to KAFKA_PEERS env)")
	fPartition = flag.Int("partition", -1, "Partition to send on")
	fTopic     = flag.String("topic", "", "Topic to send on")

	logger = log.New(os.Stderr, "producer", log.LstdFlags)
)

func flagbad(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *fPeers == "" {
		flagbad("-peers is empty\n")
	}
	if *fPartition == -1 {
		flagbad("-partition is empty\n")
	}
	if *fTopic == "" {
		flagbad("-topic is empty\n")
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "kafkaproc.producer"

	client, err := sarama.NewClient(strings.Split(*fPeers, ","), config)
	if err != nil {
		logger.Panicf("Creating sarama client: %v", err)
	}
	defer client.Close()

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		logger.Panicf("Creating sarama syncproducer: %v", err)
	}
	defer producer.Close()

	br := bufio.NewReader(os.Stdin)
	for {
		line, err := br.ReadBytes('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}
		if err != nil {
			logger.Panicf("Reading from stdin: %v", err)
		}
		line = bytes.TrimRight(line, "\n")
		message := &sarama.ProducerMessage{
			Topic:     *fTopic,
			Partition: int32(*fPartition),
			Value:     sarama.ByteEncoder(line),
		}
		part, offset, err := producer.SendMessage(message)
		if err != nil {
			logger.Panicf("Sending message: %v", err)
		}
		fmt.Printf("send (len=%d, part=%d, offset=%d)\n", len(line), part, offset)
	}
}
