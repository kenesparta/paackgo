package config

import (
	"fmt"
	"log"
	"os"
)

// Environment Get configuration from environment variables
type Environment struct{}

func (e Environment) Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Println(fmt.Sprintf("key %s environment is not set", key))
	}
	return value
}
