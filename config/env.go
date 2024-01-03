package config

import (
	"os"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("missing environment variable" + key)
	}
	return value
}
