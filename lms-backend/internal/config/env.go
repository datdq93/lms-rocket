package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// GetEnv returns the environment variable value or the provided default.
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// MustGetEnv returns the environment variable value or panics if it is empty.
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s is required", key))
	}
	return value
}

// GetEnvInt returns the integer value of an environment variable or the default when unset/invalid.
func GetEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("invalid value for %s: %v", key, err)
		} else {
			return parsed
		}
	}
	return defaultValue
}

// GetEnvDuration returns the time.Duration value of an environment variable or the default when unset/invalid.
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		parsed, err := time.ParseDuration(value)
		if err != nil {
			log.Printf("invalid duration for %s: %v", key, err)
		} else {
			return parsed
		}
	}
	return defaultValue
}
