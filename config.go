package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/gotk3/gotk3/gtk"
)

// TODO: config json here
func getConfig() (*ConfigData, error) {
	home := getEnv("XDG_HOME", "/home/aerolith/")
	config := getEnv("XDG_CONFIG", ".config")
	conf := path.Join(home, config, "/away/config.json")

	fmt.Printf("conf %v", conf)

	data := &ConfigData{}
	// var data interface{}

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

func setupFromConfig(bar *gtk.Box, config *ConfigData) error {
	for _, module := range config.Modules {
		fmt.Println("module", module.Name, module.Type)
		//TODO: replace this with a registration map
		switch module.Type {
		case "workspaces":
			mod, err := InitWorkspaces()
			if err != nil {
				fmt.Println(err)
			}
			bar.Add(mod.ToWidget())
		case "backlight":
			mod, err := InitBacklight()
			if err != nil {
				fmt.Println(err)
			}
			bar.Add(mod.ToWidget())

		case "clock":
			// mod, err := InitClock(module.timeFormat, module.dateFormat)
			// if err != nil {

			// }
			// bar.Add(mod.ToWidget())

		}
		return nil
	}
	return nil
}
