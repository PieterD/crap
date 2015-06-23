package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PieterD/kafka"
	"github.com/PieterD/kafka/internal/killchan"
	"github.com/Shopify/sarama"
)

var (
	fPeers     = flag.String("peers", os.Getenv("ZOOKEEPER_PEERS"), "List of Zookeeper peer addresses (Defaults to ZOOKEEPER_PEERS env)")
	fPartition = flag.Int("partition", -1, "Partition to receive from")
	fTopic     = flag.String("topic", "", "Topic to receive from")
	fOffset    = flag.String("offset", "newest", "newest, oldest, manual")
	fManual    = flag.Int64("manual", -1, "Set offset when -offset=manual")
	fVerbose   = flag.Bool("verbose", false, "Print message details")

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
	case "manual":
		offset = *fManual
		if offset <= -1 {
			flagbad("-manual positive integer expected")
		}
	default:
		flagbad("-offset expects newest, oldest or time\n")
	}

	kfk, err := kafka.New("kafka-consumer", logger, strings.Split(*fPeers, ","))
	if err != nil {
		logger.Panicf("Failed to create kafka client: %v", err)
	}
	defer kfk.Close()

	err = kfk.Listen(*fTopic, int32(*fPartition), offset)
	if err != nil {
		logger.Panicf("Failed to listen to partition %s:%d: %v", *fTopic, *fPartition, err)
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT)
	kc := killchan.New()
	go func() {
		<-sigchan
		kc.Kill()
	}()

	for {
		select {
		case <-kc.Chan():
			logger.Printf("Interrupt\n")
			return
		case message := <-kfk.Incoming():
			if *fVerbose {
				fmt.Printf("%s:%d (offset %d) %s\n", message.Topic, message.Partition, message.Offset, message.Val)
			} else {
				fmt.Printf("%s\n", message.Val)
			}
		}
	}
}
