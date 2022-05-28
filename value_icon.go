package main

import "github.com/gotk3/gotk3/gtk"

type ImageValue struct {
	min  float64
	max  float64
	path string
}

type ValueIcon struct {
	values []ImageValue
	image  *gtk.Image
}

func NewValueIcon(values []ImageValue) (*ValueIcon, error) {
	img, err := gtk.ImageNewFromResource("")
	if err != nil {
		return nil, err
	}
	return ValueIconFromImage(img, values)
}

func ValueIconFromImage(img *gtk.Image, values []ImageValue) (*ValueIcon, error) {
	img.SetFromFile(values[0].path)
	v := &ValueIcon{
		values: values,
		image:  img,
	}
	return v, nil
}

func (v *ValueIcon) Render(value float64) error {
	for _, val := range v.values {
		if value >= val.min && value <= val.max {
			v.image.SetFromFile(val.path)
		}
	}
	return nil
}

func (v *ValueIcon) GetWidget() gtk.IWidget {
	return v.image
}
