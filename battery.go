package main

import (
	"dlasky/away-bar/internal"
	"fmt"
	"time"

	"github.com/distatus/battery"
	"github.com/gotk3/gotk3/gtk"
)

type BatteryConfig struct {
	*internal.ModuleConfig
}

type BatteryData struct {
	Percent string
}

func InitBatteryWithConfig(cfg BatteryConfig) (gtk.IWidget, error) {

	module, err := internal.NewModuleFromConfig(cfg.ModuleConfig)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			batData, err := battery.Get(0)
			if err != nil {
				fmt.Println("[battery]", err)
			}
			data := BatteryData{
				Percent: fmt.Sprintf("%.0f", batData.Current/batData.Full*100),
			}
			module.Render(data)
			time.Sleep(60 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
