package internal

import "os"

func GetEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
