package server

import (
	"github.com/PieterD/crap/moogle/config"
	"github.com/boltdb/bolt"
)

func Run(cfg config.Config) (err error) {
	var db *bolt.DB
	db, err = bolt.Open(cfg.DatabasePath, 0600, &bolt.Options{Timeout: cfg.Timeout})
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = db.Close()
		} else {
			db.Close()
		}
	}()
	return nil
}
