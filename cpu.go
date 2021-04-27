package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/cpu"
)

func InitCPU() (gtk.IWidget, error) {

	cpuLabel, err := gtk.LabelNew("cpu")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			c, err := cpu.Percent(5*time.Second, false)
			if err != nil {
				log.Fatal("cpu fetch error")
			}

			s := fmt.Sprintf("cpu: %.0f %%", c[0])
			_, err = glib.IdleAdd(cpuLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return cpuLabel, nil
}
