package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"time"
)

func InitCPU() (gtk.IWidget, error) {

	cpuLabel, err := gtk.LabelNew("cpu")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			c, err := cpu.Percent(5*time.Second, false)

			s := fmt.Sprintf("cpu: %f%", c)
			_, err = glib.IdleAdd(cpuLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(60 * time.Second)
		}
	}()

	return cpuLabel, nil
}
