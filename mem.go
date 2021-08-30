package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/mem"
)

type MemData struct {
	UsedPercent string
}

func InitMem() (gtk.IWidget, error) {
	module, err := NewModule("memory", "{{.UsedPercent}}%", "", "./feather/box.svg")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			m, err := mem.VirtualMemory()
			if err != nil {
				log.Fatal(err)
			}
			data := MemData{
				UsedPercent: fmt.Sprintf("%.0f", m.UsedPercent),
			}
			module.Render(data)
			time.Sleep(5 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
