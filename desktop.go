package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rkoesters/xdg/desktop"
)

var desktops map[string]string = map[string]string{}

func getEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getXDGData() []string {
	out := []string{}
	home := getEnv("HOME", "~/")
	dataHome := getEnv("XDG_DATA_HOME", ".local/share")
	appLocal := filepath.Join(home, dataHome, "applications")
	out = append(out, appLocal)
	defDirs := "/usr/local/share/applications:/usr/share/applications"
	dataDirs := strings.Split(getEnv("XDG_DATA_DIRS", defDirs), ":")
	out = append(out, dataDirs...)
	return out
}

func cacheDesktops() error {
	dataDirs := getXDGData()

	for _, dir := range dataDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			fmt.Println(filepath.FromSlash(filepath.Join(dir, file.Name())))
			f, err := os.Open(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			entry, err := desktop.New(f)
			if err != nil {
				return err
			}
			desktops[entry.StartupWMClass] = entry.Icon
		}
	}
	return nil
}
