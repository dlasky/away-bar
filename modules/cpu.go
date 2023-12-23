package modules

import (
	"dlasky/away-bar/internal"
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUConfig struct {
	*internal.ModuleConfig
	UpdateInterval int `hcl:"interval"`
}

type CpuData struct {
	Percent string
	Info    cpu.InfoStat
}

func InitCPUFromConfig(cfg CPUConfig) (gtk.IWidget, error) {

	module, err := internal.NewModuleFromConfig(cfg.ModuleConfig)
	// module, err := internal.NewModule("cpu", "{{.Percent}}", "{{.Info.ModelName}}", "./feather/cpu.svg")
	if err != nil {
		return nil, err
	}

	info, err := cpu.Info()
	if err != nil {
		module.Error(err)
	}

	go func() {
		for {

			c, err := cpu.Percent(5*time.Second, false)
			if err != nil {
				module.Error(err)
			}

			data := CpuData{
				Percent: fmt.Sprintf("cpu: %.0f %%", c[0]),
				Info:    info[0],
			}
			module.Render(data)
			time.Sleep(time.Duration(cfg.UpdateInterval) * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
