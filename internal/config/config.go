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
	TMDBKey        string
}

// New config
func New() *Config {
	username := get("DATABASE_USER", "test")
	password := get("DATABASE_PASS", "test")
	url := get("DATABASE_URL", "tcp(localhost:3306)")
	databaseName := get("DATABASE_NAME", "curatedlist_test")
	googleClientID := get("GOOGLE_CLIENT_ID", "")
	tmdbKey := get("TMDB_API_KEY", "")
	log.Printf(`Generating configuration:
	- Database user: %s
	- Database url: %s
	- Database name: %s
	- Google client id set: %t
	- TMDB api key set: %t`,
		username, url, databaseName, googleClientID != "", tmdbKey != "")
	return &Config{
		DB: DBConfig{
			Username:     username,
			Password:     password,
			URL:          url,
			DatabaseName: databaseName,
		},
		GoogleClientID: googleClientID,
		TMDBKey:        tmdbKey,
	}
}

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
