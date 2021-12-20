package config

import (
	"log"
	"os"
)

// Environment Get configuration from environment variables
type Environment struct{}

func (e Environment) Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("key %s environment is not set\n", key)
	}
	return value
}
