package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	if *fConfig == "" {
		flagbad("-config is empty\n")
	}

	_, err := ParseConfig(*fConfig)
	if err != nil {
		logger.Panicf("Failed to parse config: %v", err)
	}
}
