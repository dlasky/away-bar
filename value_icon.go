package main

import "github.com/gotk3/gotk3/gtk"

type ImageValue struct {
	min  float64
	max  float64
	path string
}

type ValueIcon struct {
	values []ImageValue
	img    *gtk.Image
}

func NewValueIcon(values []ImageValue) (*gtk.Image, error) {
	img, err := gtk.ImageNewFromResource("")
	if err != nil {
		return nil, err
	}
	img.SetFromFile(values[0].path)
	return img, nil
}

func (v *ValueIcon) Render(value float64) {
	for _, val := range v.values {
		if value >= val.min && value <= val.max {
			v.img.SetFromFile(val.path)
		}
	}
}
