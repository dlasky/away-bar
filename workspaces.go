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
	wbox   *WorkspaceBox
	client sway.Client
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
	switch ev.Change {
	case "new":
		root, err := wse.client.GetTree(ctx)
		if err == nil {
			id, ok := findParentWorkspace(root, nil, &ev.Container)
			if ok {
				fmt.Println("found", id, ev.Container.ID)
				if ev.Container.AppID != nil {
					wse.wbox.AddApplication(*ev.Container.AppID, id)
				} else if ev.Container.WindowProperties != nil {
					spew.Dump(ev.Container.WindowProperties)
					wse.wbox.AddApplication(ev.Container.WindowProperties.Instance, id)
				}
			}
		}
	case "close":
	}
}

func findParentWorkspace(root *sway.Node, parent *sway.Node, target *sway.Node) (int64, bool) {
	if root.ID == target.ID {
		return parent.ID, true
	} else {
		for _, n := range root.Nodes {
			id, ok := findParentWorkspace(n, root, target)
			if ok {
				return id, ok
			}
		}
	}
	return 0, false
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
			wbox:   wbox,
			client: client,
		}
		sway.Subscribe(ctx, h, sway.EventTypeWorkspace, sway.EventTypeWindow)
	}()

	return wbox.box, nil
}
