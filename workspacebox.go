package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type WorkspaceBox struct {
	box   *gtk.Box
	list  map[int64]Workspace
	focus int64
}

type Workspace struct {
	ID    int64
	Name  string
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

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return err
	}
	label, err := gtk.LabelNew(name)
	if err != nil {
		return err
	}
	box.Add(label)
	box.ShowAll()
	ws := Workspace{
		ID,
		name,
		box,
		label,
		map[int64]*gtk.Image{},
	}

	ctx, err := box.GetStyleContext()
	if err != nil {
		return err
	}
	ctx.AddClass("workspace")

	w.box.Add(box)
	w.box.ShowAll()
	w.list[ID] = ws
	return nil
}

func (w *WorkspaceBox) Remove(ID int64) {
	if wsb, ok := w.list[ID]; ok {
		w.box.Remove(wsb.Box)
		wsb.Box.Destroy()
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
	if ws, ok := w.list[parentID]; ok {
		img, err := gtk.ImageNewFromIconName(name, gtk.ICON_SIZE_MENU)
		if err != nil {
			return err
		}
		ws.Apps[ID] = img
		ws.Box.Add(img)
		ws.Box.ShowAll()
		return nil
	}
	return nil
}

func (w *WorkspaceBox) RemoveApplication(ID int64, parentID int64) {
	if ws, ok := w.list[parentID]; ok {
		if img, ok := ws.Apps[ID]; ok {
			ws.Box.Remove(img)
			img.Destroy()
		}

	}
}
