package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("Moogle failed: %v", err)
	}
}

func run() error {
	db, err := bolt.Open(*fDatabase, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return nil
}
