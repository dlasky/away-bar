package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"time"
)

func InitMem() (gtk.IWidget, error) {
	memLabel, err := gtk.LabelNew("mem")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			m, err := mem.VirtualMemory()
			s := fmt.Sprintf("mem: %.0f %%", m.UsedPercent)
			_, err = glib.IdleAdd(memLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return memLabel, nil
}
