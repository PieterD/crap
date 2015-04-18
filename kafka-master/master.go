package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

}
