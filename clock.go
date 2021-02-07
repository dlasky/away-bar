package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"time"
)

func InitClock(format string) (gtk.IWidget, error) {

	clockLabel, err := gtk.LabelNew("clock")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			t := time.Now()
			s := t.Format(format)
			_, err = glib.IdleAdd(clockLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return clockLabel, nil
}
