package util

import (
	"os"
	"strconv"
)

func EnvString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func EnvInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if num, err := strconv.Atoi(value); err == nil {
			return num
		}
	}

	return defaultValue
}

func EnvBool(key string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if num, err := strconv.ParseBool(value); err == nil {
			return num
		}
	}

	return defaultValue
}
