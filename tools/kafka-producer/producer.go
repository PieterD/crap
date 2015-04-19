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

	"github.com/PieterD/kafka"
)

var (
	fPeers   = flag.String("peers", os.Getenv("ZOOKEEPER_PEERS"), "List of Zookeeper peer addresses (Defaults to ZOOKEEPER_PEERS env)")
	fTopic   = flag.String("topic", "", "Topic to send on")
	fVerbose = flag.Bool("verbose", false, "Print message details")

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
	if *fTopic == "" {
		flagbad("-topic is empty\n")
	}

	kfk, err := kafka.New("kafka-producer", logger, strings.Split(*fPeers, ","))
	if err != nil {
		logger.Panicf("Failed to start kafka")
	}
	defer kfk.Close()

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
		part, offset, err := kfk.Send(nil, line, *fTopic)
		if err != nil {
			logger.Panicf("Sending message: %v", err)
		}
		if *fVerbose {
			fmt.Printf("send (len=%d, part=%d, offset=%d)\n", len(line), part, offset)
		}
	}
}
