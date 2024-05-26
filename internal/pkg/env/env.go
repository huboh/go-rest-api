// Package env provides utilities for working with environment variables.
package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

// Get retrieves the value of the environment variable named by the key.
// just like `os.Getenv()`
func Get(key string) string {
	return os.Getenv(key)
}

// Load parses env files and loaded then into ENV for this process
func Load() error {
	return godotenv.Load()
}

// MustGet retrieves the value of the environment variable named by the key.
// It panics when the value is empty or unset.
func MustGet(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic(fmt.Errorf("env variable key \"%s\" not present", key))
	}

	return val
}
