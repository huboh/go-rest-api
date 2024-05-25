package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Env struct{}

// Get retrieves the value of the environment variable named by the key.
// just like `os.Getenv()`
func (e *Env) Get(key string) string {
	return os.Getenv(key)
}

// MustGet retrieves the value of the environment variable named by the key.
// It panics when the value is empty or unset.
func (e *Env) MustGet(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic(fmt.Errorf("env variable key \"%s\" not present", key))
	}

	return val
}
