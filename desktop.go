package main

import (
	"errors"
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
		if errors.Is(err, os.ErrExist) {
			continue
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".desktop" {
				f, err := os.Open(filepath.Join(dir, file.Name()))
				if errors.Is(err, os.ErrNotExist) {
					continue
				}
				entry, err := desktop.New(f)
				if err != nil {
					fmt.Println("entry", err, f.Name())
					return err
				}
				desktops[entry.StartupWMClass] = entry.Icon
			}
		}
	}
	return nil
}
