package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gotk3/gotk3/gtk"
	"github.com/joshuarubin/go-sway"
)

type WSEHandler struct {
	sway.EventHandler
	wbox *WorkspaceBox
}

func (wse WSEHandler) Workspace(ctx context.Context, ev sway.WorkspaceEvent) {
	fmt.Printf("evt: %+v\n", ev.Change)
	switch ev.Change {
	case "focus":
		wse.wbox.Focus(ev.Current.ID)
	case "init":
		wse.wbox.Add(ev.Current.Name, ev.Current.ID)
	case "empty":
		wse.wbox.Remove(ev.Current.ID)
	}

}

func (wse WSEHandler) Window(ctx context.Context, ev sway.WindowEvent) {
	spew.Dump(ev)
}

func InitWorkspaces() (gtk.IWidget, error) {
	//todo get this from app in cliapp possibly
	ctx := context.Background()
	client, err := sway.New(ctx)
	if err != nil {
		return nil, err
	}

	wbox, err := NewWorkspaceBox()
	if err != nil {
		return nil, err
	}

	list, err := client.GetWorkspaces(ctx)
	if err != nil {
		log.Fatal("Could not list workspaces, is sway running?")
	}

	for _, ws := range list {
		wbox.Add(ws.Name, ws.ID)
		if ws.Focused {
			wbox.Focus(ws.ID)
		}
	}

	go func() {
		h := WSEHandler{
			wbox: wbox,
		}
		sway.Subscribe(ctx, h, sway.EventTypeWorkspace, sway.EventTypeWindow)
	}()

	return wbox.box, nil
}
