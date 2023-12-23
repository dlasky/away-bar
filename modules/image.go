package modules

import (
	"github.com/gotk3/gotk3/gtk"
)

type ImageConfig struct {
	Source string `hcl:"source"`
}

func InitImageFromConfig(cfg *ImageConfig) (*gtk.Image, error) {
	img, err := gtk.ImageNewFromFile(cfg.Source)
	if err != nil {
		return nil, err
	}
	return img, nil
}
