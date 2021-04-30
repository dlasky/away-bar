package main

import "github.com/gotk3/gotk3/gtk"

func InitBacklight() (gtk.IWidget, error) {
	mod, err := NewModule("backlight", "", "")
	if err != nil {
		return nil, err
	}
	return mod.box, nil
}
