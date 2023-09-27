package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWT         string
}

func LoadConfig() (*Config, error) {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		return nil, fmt.Errorf("PORT environment variable not set")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PORT: %w", err)
	}

	jwt, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return nil, fmt.Errorf("JWT_KEY environment variable not set")
	}

	config := &Config{
		DatabaseURL: dbURL,
		Port:        port,
		JWT:         jwt,
	}

	return config, nil
}
