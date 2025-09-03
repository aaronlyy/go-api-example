package util

import (
	"os"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return ""
	}
	return value
}