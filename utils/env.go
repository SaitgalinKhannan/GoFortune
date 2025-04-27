package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found. Using system environment variables.")
	}
}

// MustGetEnv returns a required env variable
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Required environment variable %s not set", key)
	}
	return value
}
