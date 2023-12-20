package main

import (
	"dlasky/away-bar/internal"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type ClockConfig struct {
	Title         string               `hcl:"title,label"`
	Format        string               `hcl:"format"`
	TooltipFormat string               `hcl:"tooltip,optional"`
	Icon          *internal.IconConfig `hcl:"icon,block"`
}

func InitClockWithConfig(cfg ClockConfig) (gtk.IWidget, error) {

	type ClockData struct {
		Value   string
		Tooltip string
	}

	module, err := internal.NewModuleFromConfig(&internal.ModuleConfig{
		Name:    cfg.Title,
		Format:  "{{.Value}}",
		Tooltip: "{{.Tooltip}}",
		Icon:    cfg.Icon,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			t := time.Now()
			d := ClockData{
				Value:   t.Format(cfg.Format),
				Tooltip: t.Format(cfg.TooltipFormat),
			}

			module.Render(d)
			time.Sleep(60 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
