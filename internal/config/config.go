package config

import (
	"os"
)

// Config of the App
type Config struct {
}

// New config
func New() *Config {
	return &Config{}
}

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
