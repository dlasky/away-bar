package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func InitIcon(name string) (*gtk.Image, error) {
	icon, err := glib.IconNewForString(name)
	if err != nil {
		return nil, err
	}
	img, err := gtk.ImageNewFromGIcon(icon, gtk.ICON_SIZE_SMALL_TOOLBAR)
	if err != nil {
		return nil, err
	}
	return img, nil
}
