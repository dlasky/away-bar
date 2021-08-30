package main

import (
	"fmt"
	"log"
	"time"

	"github.com/distatus/battery"
	"github.com/gotk3/gotk3/gtk"
)

type BatteryData struct {
	Percent string
}

func InitBattery() (gtk.IWidget, error) {

	module, err := NewModule("battery", "{{.Percent}}", "", "./feather/battery.svg")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			batData, err := battery.Get(0)
			if err != nil {
				log.Fatal(err)
			}
			data := BatteryData{
				Percent: fmt.Sprintf("bat: %.0f %%", batData.Current/batData.Full*100),
			}
			module.Render(data)
			time.Sleep(60 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
