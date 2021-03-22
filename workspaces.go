package main

import (
	"bytes"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"go.i3wm.org/i3/v4"
	"log"
	"os/exec"
)

type wsLabel struct {
	label *gtk.Button
	ws    i3.Workspace
}

func InitWorkspaces() (gtk.IWidget, error) {

	//i3 overrides to work with sway
	i3.SocketPathHook = func() (string, error) {
		out, err := exec.Command("sway", "--get-socketpath").CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("getting sway socketpath: %v (output: %s)", err, out)
		}
		return string(out), nil
	}

	i3.IsRunningHook = func() bool {
		out, err := exec.Command("pgrep", "-c", "sway\\$").CombinedOutput()
		if err != nil {
			log.Printf("sway running: %v (output: %s)", err, out)
		}
		return bytes.Compare(out, []byte("1")) == 0
	}

	wsList, err := i3.GetWorkspaces()
	if err != nil {
		return nil, err
	}

	b, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return nil, err
	}

	var wsWidgets = []wsLabel{}

	for i, w := range wsList {
		l, err := gtk.ButtonNew()
		l.SetLabel(w.Name)
		if err != nil {
			return nil, err
		}
		i := i
		l.Connect("clicked", func() {
			i3.RunCommand(fmt.Sprintf("workspace %v", i+1))
		})
		wsWidgets = append(wsWidgets, wsLabel{l, w})
		b.Add(l)
	}

	go func() {

		recv := i3.Subscribe(i3.WorkspaceEventType)
		for recv.Next() {
			ev := recv.Event().(*i3.WorkspaceEvent)

			for _, wsl := range wsWidgets {
				if int64(ev.Old.ID) == int64(wsl.ws.ID) {

					_, err := glib.IdleAdd(wsl.label.SetLabel, ev.Old.Name)
					if err != nil {
						log.Fatal("ui error")
					}
				}
				if int64(ev.Current.ID) == int64(wsl.ws.ID) {
					glib.IdleAdd(wsl.label.SetLabel, fmt.Sprintf("[%v]", ev.Current.Name))
					if err != nil {
						log.Fatal("ui error")
					}
				}
			}
		}
		log.Fatal(recv.Close())
	}()

	return b, nil
}
