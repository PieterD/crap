package main

import (
	"flag"
	"log"
	"time"

	"github.com/PieterD/crap/moogle/config"
	"github.com/PieterD/crap/moogle/server"
)

var fDatabasePath = flag.String("database", "/tmp/boltdb", "Path to database")

func main() {
	flag.Parse()

	err := server.Run(config.Config{
		DatabasePath: *fDatabasePath,
		Timeout:      time.Second,
	})
	if err != nil {
		log.Printf("Moogle failed: %v", err)
	}
}
