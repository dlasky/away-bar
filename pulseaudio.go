package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/lawl/pulseaudio"
)

type VolumeData struct {
	Volume string
	Muted  string
}

func InitPulseAudio() (gtk.IWidget, error) {

	module, err := NewModule("volume", "{{.Volume}} %", "", "./feather/speaker.svg")
	if err != nil {
		return nil, err
	}

	client, err := pulseaudio.NewClient()
	if err != nil {
		module.error(err)
	}

	go func() {
		cha, _ := client.Updates()
		vol, err := client.Volume()
		if err != nil {
			module.error(err)
		}
		data := VolumeData{Volume: fmt.Sprintf("%.0f", vol*100)}
		module.Render(data)
		go func() {
			for {
				<-cha
				vol, err := client.Volume()
				if err != nil {
					module.error(err)
				}
				data := VolumeData{
					Volume: fmt.Sprintf("%.0f", vol*100),
				}
				module.Render(data)
			}
		}()
	}()

	return module.GetWidget(), nil
}
