package main

import (
	"dlasky/away-bar/internal"
	"dlasky/away-bar/modules"
	"fmt"
	"log"
	"path"

	"github.com/gotk3/gotk3/gtk"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func getConfig() (*Config, error) {

	//TODO: allow an explicit pathing here as well via flags

	user := internal.GetEnv("USER", "")
	home := internal.GetEnv("XDG_HOME", "/home/"+user)
	cfg := internal.GetEnv("XDG_CONFIG", ".config")
	conf := path.Join(home, cfg, "/awaybar/config.hcl")

	fmt.Printf("conf %v", conf)

	var config Config
	err := hclsimple.DecodeFile(conf, nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)
	return &config, err
}

func setupFromConfig(bar *gtk.Box, config *Config) error {

	if config.Bar.Clock != nil {
		w, err := modules.InitClockWithConfig(*config.Bar.Clock)
		if err != nil {
			fmt.Println(err)
		}
		bar.Add(w.ToWidget())
	}

	return nil
}
