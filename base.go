package main

import "github.com/gotk3/gotk3/gtk"

func initBase(name string, defaultValue string) (*gtk.Label, error) {
	lbl, err := gtk.LabelNew(defaultValue)
	if err != nil {
		return nil, err
	}
	lbl.SetName(name)
	return lbl, nil
}
