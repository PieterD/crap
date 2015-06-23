package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PieterD/crap/kafka"
	"github.com/PieterD/crap/kafka/internal/killchan"
)

var (
	fPeers       = flag.String("peers", os.Getenv("ZOOKEEPER_PEERS"), "List of Zookeeper peer addresses (Defaults to ZOOKEEPER_PEERS env)")
	fPartition   = flag.Int("partition", -1, "Partition to receive from")
	fSource      = flag.String("src", "", "Topic to receive from")
	fDestination = flag.String("dst", "", "Topic to send to")
	fVerbose     = flag.Bool("verbose", false, "Be more verbose")

	logger = log.New(os.Stderr, "processor", log.LstdFlags)
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
	if *fSource == "" {
		flagbad("-src is empty\n")
	}
	if *fDestination == "" {
		flagbad("-dst is empty\n")
	}

	kfk, err := kafka.New("kafka-processor", logger, strings.Split(*fPeers, ","))
	if err != nil {
		logger.Panicf("Failed to start kafka: %v\n", err)
	}
	defer kfk.Close()

	err = kfk.ListenNewest(*fSource, int32(*fPartition))
	if err != nil {
		logger.Panicf("Failed to listen to partition %s:%d: %v\n", *fSource, *fPartition)
	}
	defer kfk.Unlisten(*fSource, int32(*fPartition))

	kill := killchan.New()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT)
	go func() {
		<-sigchan
		kill.Kill()
		<-sigchan
		panic("second SIGINT")
	}()

	for {
		var msg kafka.Message

		select {
		case <-kill.Chan():
			return
		case msg = <-kfk.Incoming():
		}

		// PROCESSING
		newVal := bytes.ToUpper(msg.Val)
		// PROCESSING

		select {
		case <-kill.Chan():
			return
		case trans := <-kfk.Outgoing():
			partition, offset, err := trans.Send(msg.Key, newVal, *fDestination)
			if err != nil {
				logger.Panicf("Message send failed to %s: %v", *fDestination, err)
			}
			if *fVerbose {
				fmt.Printf("processed %s:%d(%d) -> %s:%d(%d)\n", msg.Topic, msg.Partition, msg.Offset, *fDestination, partition, offset)
			}
		}
	}
}
