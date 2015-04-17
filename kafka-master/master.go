package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	fZooPeers = flag.String("zoopeers", os.Getenv("ZOOKEEPER_PEERS"), "List of Zookeeper peer addresses (Defaults to ZOOKEEPER_PEERS env)")

	logger = log.New(os.Stderr, "master", log.LstdFlags)
)

func flagbad(f string, i ...interface{}) {
	fmt.Fprintf(os.Stderr, f, i...)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *fZooPeers == "" {
		flagbad("No zookeeper peers provided; use -zoopeers or the ZOOKEEPER_PEERS environment variable")
	}
	//func Connect(servers []string, sessionTimeout time.Duration) (*Conn, <-chan Event, error)
	zooConn, _, err := zk.Connect(strings.Split(*fZooPeers, ","), time.Second)
	if err != nil {
		logger.Panicf("Failed to connect to zookeeper: %v", err)
	}

	brokers := GetBrokers(zooConn)
	for _, broker := range brokers {
		fmt.Printf("broker %3d, addr: %s\n", broker.Id, broker.Addr())
	}
}
