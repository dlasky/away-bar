package workspace

import (
	"context"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/joshuarubin/go-sway"
)

type WSEHandler struct {
	sway.EventHandler
	wbox   *WorkspaceBox
	client sway.Client
}

var tree *sway.Node

func (wse WSEHandler) Workspace(ctx context.Context, ev sway.WorkspaceEvent) {
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
	fmt.Println("change", ev.Change)
	switch ev.Change {
	case "new":
		root, err := wse.client.GetTree(ctx)
		if err != nil {
			fmt.Println("tree error", err)
		}
		tree = root
		list := traverse(root, 0)
		for _, app := range list {
			if ev.Container.ID == app.ID {
				wse.wbox.AddApplication(app.name, app.ID, app.parentID)
			}
		}
	case "close":
		list := traverse(tree, 0)
		for _, app := range list {
			fmt.Println("rm", ev.Container.ID, app.ID, app.parentID)
			if ev.Container.ID == app.ID {
				wse.wbox.RemoveApplication(ev.Container.ID, app.parentID)
			}
		}
	}
}

func getName(node *sway.Node) string {
	if node.AppID != nil {
		fmt.Println(node.AppID)
		return *node.AppID
	} else if node.WindowProperties != nil {
		name, ok := desktops[node.WindowProperties.Class]
		if ok {
			return name
		}
		return node.WindowProperties.Instance
	}
	fmt.Println("didnt find name", node.AppID)
	return ""
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

	root, err := client.GetTree(ctx)
	if err != nil {
		log.Fatal("Could not get sway window tree is sway running?")
	}
	apps := traverse(root, 0)
	for _, app := range apps {
		wbox.AddApplication(app.name, app.ID, app.parentID)
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

type WorkspaceLeaf struct {
	name     string
	ID       int64
	parentID int64
}

func traverse(n *sway.Node, wsID int64) []WorkspaceLeaf {
	if n.Type == "workspace" {
		wsID = n.ID
	}
	if len(n.Nodes) == 0 {
		return []WorkspaceLeaf{
			{
				name:     getName(n),
				ID:       n.ID,
				parentID: wsID,
			},
		}
	}
	var output = []WorkspaceLeaf{}
	for _, node := range n.Nodes {
		output = append(output, traverse(node, wsID)...)
	}
	return output
}
