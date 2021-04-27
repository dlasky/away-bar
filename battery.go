package main

import (
	"fmt"
	"log"
	"time"

	"github.com/distatus/battery"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func InitBattery() (gtk.IWidget, error) {
	batLabel, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			batData, err := battery.Get(0)
			if err != nil {
				log.Fatal(err)
			}
			s := fmt.Sprintf("bat: %.0f %%", batData.Current/batData.Full*100)
			_, err = glib.IdleAdd(batLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(60 * time.Second)
		}
	}()

	return batLabel, nil
}
