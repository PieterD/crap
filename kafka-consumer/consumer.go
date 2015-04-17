package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	fPeers     = flag.String("peers", os.Getenv("KAFKA_PEERS"), "List of Kafka peer addresses (Defaults to KAFKA_PEERS env)")
	fPartition = flag.Int("partition", -1, "Partition to send on")
	fTopic     = flag.String("topic", "", "Topic to send on")
	fOffset    = flag.String("offset", "newest", "newest, oldest")
	fVerbose   = flag.Bool("verbose", false, "Print message details")
	//fTime      = flag.String("time", "", "Offset time (if -offset=time): yyyy-mm-ddThh:mm:ss.nnnnnnnnn")

	logger = log.New(os.Stderr, "consumer", log.LstdFlags)
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
	var offset int64
	switch *fOffset {
	case "newest":
		offset = sarama.OffsetNewest
	case "oldest":
		offset = sarama.OffsetOldest
		/*
			case "time":
				t, err := time.Parse("2006-01-02T15:04:05.999999999", *fTime)
				if err != nil {
					flagbad("-offset expects newest, oldest or time\n")
				}
		*/
	default:
		flagbad("-offset expects newest, oldest or time\n")
	}

	config := sarama.NewConfig()
	config.ClientID = "kafkaproc.consumer"

	client, err := sarama.NewClient(strings.Split(*fPeers, ","), config)
	if err != nil {
		logger.Panicf("Creating sarama client: %v", err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		logger.Panicf("Creating sarama consumer: %v", err)
	}
	defer consumer.Close()

	partitionconsumer, err := consumer.ConsumePartition(*fTopic, int32(*fPartition), offset)
	if err != nil {
		logger.Panicf("Creating partition consumer: %v", err)
	}
	defer partitionconsumer.Close()

	for message := range partitionconsumer.Messages() {
		if *fVerbose {
			fmt.Printf("(offset=%d) %s\n", message.Offset, message.Value)
		} else {
			fmt.Printf("%s\n", message.Value)
		}
	}
}
