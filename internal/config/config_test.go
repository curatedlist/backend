package config_test

import (
	"backend/internal/config"
	"os"
	"testing"
)

func TestLookupEnv(t *testing.T) {
	const username = "test"
	defer os.Unsetenv("DATABASE_USER")
	err := os.Setenv("DATABASE_USER", "alice")
	if err != nil {
		t.Fatalf("failed to set environment variable username to test")
	}

	conf := config.New()
	if conf.DB.Username != "alice" {
		t.Errorf("failed to read environment variable")
	}
}
