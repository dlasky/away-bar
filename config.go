package main

import (
	"encoding/json"
	"os"
	"path"

	"github.com/gotk3/gotk3/gtk"
)

//TODO: config json here
func getConfig() (*ConfigData, error) {
	home := getEnv("XDG_HOME", "~/")
	config := getEnv("XDG_CONFIG", ".config")
	conf := path.Join(home, config, "/away/config.json")

	data := &ConfigData{}

	byt, err := os.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byt, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setupFromConfig(bar *gtk.Box) error {
	data, err := getConfig()
	if err != nil {
		return err
	}
	for _, module := range data.Modules {
		//TODO: replace this with a registration map
		switch module.Type {

		}

	}
	return nil
}
