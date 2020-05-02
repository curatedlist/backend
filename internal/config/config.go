package config

import (
	"os"
)

//DBConfig is the config of the database
type DBConfig struct {
	Username string
	Password string
	URL      string
}

// Config of the App
type Config struct {
	DB DBConfig
}

// New config
func New() *Config {
	return &Config{
		DB: DBConfig{
			Username: get("DATABASE_USER", "test"),
			Password: get("DATABASE_PASS", "test"),
			URL:      get("DATABASE_URL", "tcp(localhost:3306)"),
		},
	}
}

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
