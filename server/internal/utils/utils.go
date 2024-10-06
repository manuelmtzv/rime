package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvOrThrow(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Missing environment variable: %s", key))
	}
	return value
}

func GetEnvAsIntOrThrow(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Missing environment variable: %s", key))
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Invalid integer for %s: %v", key, err))
	}

	return intValue
}
