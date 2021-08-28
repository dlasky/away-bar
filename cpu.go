package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuData struct {
	Percent string
	Info    cpu.InfoStat
}

func InitCPU() (gtk.IWidget, error) {

	module, err := NewModule("cpu", "{{.Percent}}", "{{.Info.ModelName}}", "./feather/cpu.svg")
	if err != nil {
		return nil, err
	}

	info, err := cpu.Info()
	spew.Dump(info)
	if err != nil {
		module.error(err)
	}

	go func() {
		for {

			c, err := cpu.Percent(5*time.Second, false)
			if err != nil {
				module.error(err)
			}

			data := CpuData{
				Percent: fmt.Sprintf("cpu: %.0f %%", c[0]),
				Info:    info[0],
			}
			module.Render(data)
			time.Sleep(5 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
