package config

import (
	"time"

	"github.com/PieterD/crap/moogle/input"
)

type Config struct {
	DatabasePath string
	Timeout      time.Duration
	Inputs       []input.Input
}
