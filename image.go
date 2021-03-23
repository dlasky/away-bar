package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func InitImage(src string) *gtk.Image {
	img, err := gtk.ImageNewFromFile(src)
	if err != nil {
		log.Fatalf("image error %v", err)
	}
	return img
}
