package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PieterD/kafka"
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

	kfk, err := kafka.New("kafka-processor", logger, cfg.Zookeeper.Peer)
	if err != nil {
		logger.Panicf("Failed to start kafka")
	}
	defer kfk.Close()

	th, err := newThreadHandler(kfk, cfg)
	if err != nil {
		logger.Panicf("Failed to start thread handler")
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT)
	go func() {
		<-sigchan
		th.close()
	}()
	th.wait()
}
