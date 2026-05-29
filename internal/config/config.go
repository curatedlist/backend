package config

import (
	"log"
	"os"
)

//DBConfig is the config of the database
type DBConfig struct {
	Username     string
	Password     string
	URL          string
	DatabaseName string
}

// Config of the App
type Config struct {
	DB             DBConfig
	GoogleClientID string
}

// New config
func New() *Config {
	username := get("DATABASE_USER", "test")
	password := get("DATABASE_PASS", "test")
	url := get("DATABASE_URL", "tcp(localhost:3306)")
	databaseName := get("DATABASE_NAME", "curatedlist_test")
	googleClientID := get("GOOGLE_CLIENT_ID", "")
	log.Printf(`Generating configuration:
	- Database user: %s
	- Database url: %s
	- Database name: %s
	- Google client id set: %t`,
		username, url, databaseName, googleClientID != "")
	return &Config{
		DB: DBConfig{
			Username:     username,
			Password:     password,
			URL:          url,
			DatabaseName: databaseName,
		},
		GoogleClientID: googleClientID,
	}
}

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
