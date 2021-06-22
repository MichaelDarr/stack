package config

import (
	"log"
	"os"
	"strconv"
)

// GetBoolEnv returns an boolean environment variable with a fallback if not found or invalid.
func GetBoolEnv(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if boolVal, err := strconv.ParseBool(val); err == nil {
			return boolVal
		}
	}
	return fallback
}

// GetRequiredEnv returns an environment variable or panics if not present.
func GetRequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("required environment variable not found: %s", key)
	}
	return val
}

// GetRequiredBoolEnv returns an environment variable or panics if not present or not a boolean.
func GetRequiredBoolEnv(key string) bool {
	val := GetRequiredEnv(key)
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("required environment variable could not be converted to boolean: %s", key)
	}
	return boolVal
}

// GetRequiredIntEnv returns an environment variable or panics if not present or not an integer.
func GetRequiredIntEnv(key string) int {
	val := GetRequiredEnv(key)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("required environment variable could not be converted to integer: %s", key)
	}
	return intVal
}
