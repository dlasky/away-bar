package main

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type ClockConfig struct {
	Format        string `hcl:"format"`
	TooltipFormat string `hcl:"tooltip,optional"`
}

func InitClock(format string, tooltipFormat string) (gtk.IWidget, error) {

	type ClockData struct {
		Value   string
		Tooltip string
	}

	module, err := NewModule("clock", "{{.Value}}", "{{.Tooltip}}", "./feather/clock.svg")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			t := time.Now()
			d := ClockData{
				Value:   t.Format(format),
				Tooltip: t.Format(tooltipFormat),
			}

			module.Render(d)
			time.Sleep(60 * time.Second)
		}
	}()

	return module.GetWidget(), nil
}
