package utils

import (
	"os"
)

// GetEnv looks for an existing value in the
// environemnt and returns it or the fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}