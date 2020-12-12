package utils

import (
	"os"
)

// GetEnv looks for an existing value in
// the enviroment vars and returns it.
// if vars not existing, fallback will be execute.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}