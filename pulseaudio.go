package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/lawl/pulseaudio"
)

func InitPulseAudio() (gtk.IWidget, error) {
	client, err := pulseaudio.NewClient()
	if err != nil {
		return nil, err
	}

	volLabel, err := gtk.LabelNew("volume")
	if err != nil {
		return nil, err
	}

	go func() {
		cha, _ := client.Updates()
		vol, err := client.Volume()
		if err != nil {
			log.Fatal("pulse issue")
		}
		s := fmt.Sprintf("vol: %.0f %%", vol*100)
		_, err = glib.IdleAdd(volLabel.SetText, s)
		if err != nil {
			log.Fatal("IdleAdd() failed:", err)
		}
		go func() {
			for {
				<-cha
				vol, err := client.Volume()
				if err != nil {
					log.Fatal("pulse issue")
				}
				s := fmt.Sprintf("vol: %.0f %%", vol*100)
				_, err = glib.IdleAdd(volLabel.SetText, s)
				if err != nil {
					log.Fatal("IdleAdd() failed:", err)
				}
			}
		}()
	}()

	// go func() {
	// 	for {
	// 		vol, err := client.Volume()
	// 		if err != nil {
	// 			log.Fatal("pulse issue")
	// 		}
	// 		s := fmt.Sprintf("vol: %.0f %%", vol*100)
	// 		_, err = glib.IdleAdd(volLabel.SetText, s)
	// 		if err != nil {
	// 			log.Fatal("IdleAdd() failed:", err)
	// 		}
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }()

	return volLabel, nil
}
