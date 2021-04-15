package config

import (
	"log"
	"os"
	"strconv"
)

// MustGet will return an environment variable or panic if it is not present
func GetRequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("required environment variable not found: %s", key)
	}
	return val
}

func GetRequiredBoolEnv(key string) bool {
	val := GetRequiredEnv(key)
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("required environment variable could not be converted to boolean: %s", key)
	}
	return boolVal
}

// GetRequiredIntEnv will return an environment variable or panic if it is not present or not an integer
func GetRequiredIntEnv(key string) int {
	val := GetRequiredEnv(key)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("required environment variable could not be converted to integer: %s", key)
	}
	return intVal
}
