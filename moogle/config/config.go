package config

import "time"

type Config struct {
	DatabasePath string
	Timeout      time.Duration
}
