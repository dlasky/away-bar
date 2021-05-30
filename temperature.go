package main

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/host"
)

func InitTemp() (gtk.IWidget, error) {

	module, err := NewModule("temperature", "temp: {{(index . 0).Temperature}}C", "./feather/thermometer.svg")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			temps, err := host.SensorsTemperatures()
			if err != nil {
				log.Fatal("error getting temps")
			}
			module.Render(temps)
			time.Sleep(5 * time.Second)
		}
	}()
	return module.box, nil
}
