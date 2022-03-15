package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type ImageRenderable struct {
	image *gtk.Image
}

func InitImage(src string) (*ImageRenderable, error) {
	img, err := gtk.ImageNewFromFile(src)
	if err != nil {
		return nil, err
	}
	return &ImageRenderable{
		image: img,
	}, nil
}

func (i *ImageRenderable) Render(path interface{}) error {
	i.image.SetFromFile(fmt.Sprintf("%v", path))
	return nil
}

func (i *ImageRenderable) GetWidget() *gtk.Image {
	return i.image
}
