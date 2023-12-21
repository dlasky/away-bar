package main

import (
	"dlasky/away-bar/internal"
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/mem"
)

type MemConfig struct {
	*internal.ModuleConfig
	UpdateInterval int `hcl:"interval"`
}
type MemData struct {
	UsedPercent string
}

func InitMemFromConfig(cfg MemConfig) (gtk.IWidget, error) {

	module, err := internal.NewModuleFromConfig(cfg.ModuleConfig)
	// module, err := NewModule("memory", "{{.UsedPercent}}%", "", "./feather/box.svg")
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
			time.Sleep(time.Duration(cfg.UpdateInterval) * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
