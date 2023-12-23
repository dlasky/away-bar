package modules

import (
	"dlasky/away-bar/internal"
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/lawl/pulseaudio"
)

type PulseAudioConfig struct {
	*internal.ModuleConfig
}

type VolumeData struct {
	Volume string
	Muted  string
}

func InitPulseAudioFromConfig(cfg PulseAudioConfig) (gtk.IWidget, error) {

	module, err := internal.NewModuleFromConfig(cfg.ModuleConfig)
	// module, err := NewModule("volume", "{{.Volume}} %", "", "./feather/speaker.svg")
	if err != nil {
		return nil, err
	}

	client, err := pulseaudio.NewClient()
	if err != nil {
		module.Error(err)
	}

	go func() {
		cha, _ := client.Updates()
		vol, err := client.Volume()
		if err != nil {
			module.Error(err)
		}
		data := VolumeData{Volume: fmt.Sprintf("%.0f", vol*100)}
		module.Render(data)
		go func() {
			for {
				<-cha
				vol, err := client.Volume()
				if err != nil {
					module.Error(err)
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
