package internal

import (
	"os"
	"path"
)

func GetEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetConfigPath(file string) string {
	user := GetEnv("USER", "user")
	home := GetEnv("XDG_HOME", "/home/"+user)
	cfg := GetEnv("XDG_CONFIG", ".config")
	conf := path.Join(home, cfg, "/awaybar/", file)
	return conf
}
