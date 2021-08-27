package main

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func InitClock(format string) (gtk.IWidget, error) {

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

	clockLabel, err := gtk.LabelNew("clock")
	if err != nil {
		return nil, err
	}

	box.SetTooltipText("clock face")

	box.Add(clockLabel)
	box.ShowAll()

	go func() {
		for {
			t := time.Now()
			s := t.Format(format)
			_, err = glib.IdleAdd(clockLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			time.Sleep(60 * time.Second)
		}
	}()

	return box, nil
}
