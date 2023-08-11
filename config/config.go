package config

import (
	"os"
)

type Config struct {
	IngestFilePath string
	DBHost         string
	DBPort         string
	DBUsername     string
	DBPassword     string
	DBname         string
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetConfig() *Config {
	return &Config{
		IngestFilePath: getenv("INGEST_FILE_PATH", "ingest/ports.json"),
		DBHost:         getenv("DB_HOST", "localhost"),
		DBPort:         getenv("DB_PORT", "5432"),
		DBUsername:     getenv("DB_USERNAME", "postgres"),
		DBPassword:     getenv("DB_PASSWORD", "postgres"),
		DBname:         getenv("DB_NAME", "dev"),
	}
}
