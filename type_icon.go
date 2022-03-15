package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type TypeIcon struct {
	types map[string]string
	image *gtk.Image
}

func NewTypeIcon(config map[string]string, initial string) (*TypeIcon, error) {
	img, err := gtk.ImageNewFromResource("")
	if err != nil {
		return nil, err
	}
	val, ok := config[initial]
	if ok {
		img.SetFromFile(val)
	}
	t := TypeIcon{
		types: config,
		image: img,
	}
	return &t, nil
}

func (t *TypeIcon) Render(value interface{}) error {
	val, ok := t.types[fmt.Sprintf("%v", value)]
	if ok {
		t.image.SetFromFile(val)
	}
	return nil
}

func (t *TypeIcon) GetWidget() gtk.IWidget {
	return t.image
}
