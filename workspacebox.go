package main

import (
	"context"
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/joshuarubin/go-sway"
)

type WorkspaceBox struct {
	box   *gtk.Box
	list  map[int64]Workspace
	focus int64
}

type Workspace struct {
	ID    int64
	Name  string
	EBox  *gtk.EventBox
	Box   *gtk.Box
	Label *gtk.Label
	Apps  map[int64]*gtk.Image
}

func NewWorkspaceBox() (*WorkspaceBox, error) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		return nil, err
	}
	return &WorkspaceBox{
		box:  box,
		list: map[int64]Workspace{},
	}, nil
}

func (w *WorkspaceBox) Add(name string, ID int64) error {

	fmt.Println("Add", name, ID)

	ebox, err := gtk.EventBoxNew()
	if err != nil {
		return err
	}
	ebox.Connect("button-release-event", func() {
		bg := context.Background()
		client, err := sway.New(bg)
		if err != nil {
			fmt.Println(err)
		}
		client.RunCommand(bg, "workspace "+name)
	})

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return err
	}
	ebox.Add(box)
	label, err := gtk.LabelNew(name)
	if err != nil {
		return err
	}
	box.Add(label)
	box.ShowAll()
	ws := Workspace{
		ID,
		name,
		ebox,
		box,
		label,
		map[int64]*gtk.Image{},
	}

	ctx, err := box.GetStyleContext()
	if err != nil {
		return err
	}
	ctx.AddClass("workspace")

	w.box.Add(ebox)
	w.box.ShowAll()
	w.list[ID] = ws
	return nil
}

func (w *WorkspaceBox) Remove(ID int64) {

	fmt.Println("remove", ID)

	if wsb, ok := w.list[ID]; ok {
		w.box.Remove(wsb.EBox)
		wsb.EBox.Destroy()
		wsb.Box.Destroy()
		//todo clear apps as well?
		delete(w.list, ID)
	}
}

func (w *WorkspaceBox) Focus(ID int64) error {

	if ws, ok := w.list[w.focus]; ok {
		err := setButtonFocus(ws.Box, false)
		if err != nil {
			return err
		}
	}

	if wsb, ok := w.list[ID]; ok {
		w.focus = ID
		err := setButtonFocus(wsb.Box, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func setButtonFocus(btn *gtk.Box, focus bool) error {
	const fc = "focused"
	ctx, err := btn.GetStyleContext()
	if err != nil {
		return err
	}
	if focus {
		ctx.AddClass(fc)
	} else {
		ctx.RemoveClass(fc)
	}
	return nil
}

func (w *WorkspaceBox) AddApplication(name string, ID int64, parentID int64) error {
	fmt.Println("AddApplication", name, ID, parentID)
	if ws, ok := w.list[parentID]; ok {
		img, err := gtk.ImageNewFromIconName(name, gtk.ICON_SIZE_MENU)
		if err != nil {
			return err
		}
		img.Connect("clicked", func() {
			fmt.Println("clicked")
		})
		ws.Apps[ID] = img
		ws.Box.Add(img)
		ws.Box.ShowAll()
		return nil
	} else {
		fmt.Println("no parent for", ID, parentID)
		for k := range w.list {
			fmt.Println(k)
		}
	}
	return nil
}

func (w *WorkspaceBox) RemoveApplication(ID int64, parentID int64) {
	fmt.Println("removeApplication", ID, parentID)
	if ws, ok := w.list[parentID]; ok {
		if img, ok := ws.Apps[ID]; ok {
			ws.Box.Remove(img)
			img.Destroy()
		}

	}
}
