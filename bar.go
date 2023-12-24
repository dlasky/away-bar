package main

import "github.com/gotk3/gotk3/gtk"

type Orientation string

const (
	Top    Orientation = "top"
	Left   Orientation = "left"
	Right  Orientation = "right"
	Bottom Orientation = "bottom"
)

type BarConfig struct {
	Position Orientation `hcl:"orientation"`
	Display  string      `hcl:"display"`
}

type Bar struct {
	Bar *gtk.Box
}

func NewBarFromConfig(cfg *BarConfig) (*Bar, error) {
	bar, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}
	return &Bar{Bar: bar}, nil
}
