package utils

import "os"

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}

func IsProd() bool {
	if env := os.Getenv("GO_ENV"); (env == "prod") || (env == "production") {
		return true
	}

	return false
}
