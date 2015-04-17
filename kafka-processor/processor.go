package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

var (
	fConfig = flag.String("config", "", "xml file configuring processing streams")

	logger = log.New(os.Stderr, "processor", log.LstdFlags)
)

func flagbad(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *fConfig == "" {
		flagbad("-config is empty\n")
	}

	cfg, err := ParseConfig(*fConfig)
	if err != nil {
		logger.Panicf("Failed to parse config: %v", err)
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = "kafkaproc.processor"

	client, err := sarama.NewClient(cfg.Kafka.Peer, config)
	if err != nil {
		logger.Panicf("Creating sarama client: %v", err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		logger.Panicf("Creating sarama consumer: %v", err)
	}
	defer consumer.Close()

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		logger.Panicf("Creating sarama syncproducer: %v", err)
	}
	defer producer.Close()

	cl := NewConsumerList()

	for _, stream := range cfg.Stream {
		for _, partition := range stream.Partition {
			partitionconsumer, err := consumer.ConsumePartition(stream.TopicSrc, partition, sarama.OffsetNewest)
			if err != nil {
				logger.Printf("Failed to create partition consumer %s:%d: %v", stream.TopicSrc, partition, err)
				continue
			}
			pc := &ProcConsumer{
				consumer:  partitionconsumer,
				producer:  producer,
				src:       stream.TopicSrc,
				dst:       stream.TopicDst,
				partition: partition,
				function:  stream.Function,
				killchan:  make(chan struct{}),
			}
			fmt.Printf("Starting %s:%d -> %s\n", stream.TopicSrc, partition, stream.TopicDst)
			cl.Add(pc)
		}
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT)
	go func() {
		<-sigchan
		fmt.Printf("SIGINT\n")
		cl.Close()
	}()
	fmt.Printf("wait\n")
	<-cl.killchan
	fmt.Printf("waited\n")
}
