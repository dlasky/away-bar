package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

func cacheDesktops() {
	dataDirs := getXDGData()

	for _, dir := range dataDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("desktop directory error:", err)
			continue
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".desktop" {
				f, err := os.Open(filepath.Join(dir, file.Name()))
				if err != nil {
					fmt.Println("desktop file error:", err)
					continue
				}
				entry, err := desktop.New(f)
				if err != nil {
					fmt.Println("desktop parse error:", err, file.Name())
					continue
				}
				desktops[entry.StartupWMClass] = entry.Icon
				desktops[entry.Exec] = entry.Icon
				desktops[entry.Name] = entry.Icon
			}
		}
	}
	spew.Dump(desktops)
}
