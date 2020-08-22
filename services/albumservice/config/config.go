package config

import (
	"os"
)

var fallback = map[string]string{
	"DB_PATH":   "./album.db",
	"ROOT_PATH": "/var/db",
}

// GetENV godoc
// Return the env value of input key
// @params { string } , key of env
func GetENV(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback[key]
	}
	return val
}
