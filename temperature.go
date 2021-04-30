package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/host"
)

func InitTemp() (gtk.IWidget, error) {
	tempLabel, err := initBase("temperature", "")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			temps, err := host.SensorsTemperatures()
			if err != nil {
				log.Fatal("error getting temps")
			}

			s := fmt.Sprintf("temp: %.0f %%", temps[0].Temperature)
			_, err = glib.IdleAdd(tempLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return tempLabel, nil
}
