package main

import (
	"dlasky/away-bar/internal"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/host"
)

type TempConfig struct {
	*internal.ModuleConfig
	//TODO: freq
}

func InitTempFromConfig(cfg TempConfig) (gtk.IWidget, error) {

	// module, err := NewModule("temperature", "temp: {{(index . 0).Temperature}}°C", "{{(index . 0).Temperature}}°C", "./feather/thermometer.svg")
	module, err := internal.NewModuleFromConfig(cfg.ModuleConfig)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			temps, err := host.SensorsTemperatures()
			if err != nil {
				module.Error(err)
			}
			module.Render(temps)
			time.Sleep(5 * time.Second)
		}
	}()
	return module.GetWidget(), nil
}
