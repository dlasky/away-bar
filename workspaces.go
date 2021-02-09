package main

import (
	"bytes"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"go.i3wm.org/i3/v4"
	"log"
	"os/exec"
)

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

	var wsWidgets = []*gtk.Label{}

	for _, w := range wsList {
		l, err := gtk.LabelNew(w.Name)
		if err != nil {
			return nil, err
		}
		wsWidgets = append(wsWidgets, l)
		b.Add(l)
	}

	go func() {

		recv := i3.Subscribe(i3.WorkspaceEventType)
		for recv.Next() {
			ev := recv.Event().(*i3.WorkspaceEvent)
			log.Printf("ws: %v", ev.Old)
		}
		log.Fatal(recv.Close())
	}()

	return b, nil
}
