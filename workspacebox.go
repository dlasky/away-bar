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
	ID     int64
	Name   string
	Button *gtk.Button
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

	btn, err := gtk.ButtonNew()
	if err != nil {
		return err
	}
	btn.SetLabel(name)
	ws := Workspace{
		ID,
		name,
		btn,
	}

	ctx, err := btn.GetStyleContext()
	if err != nil {
		return err
	}
	ctx.AddClass("workspace")

	w.box.Add(btn)
	w.list[ID] = ws
	return nil
}

func (w *WorkspaceBox) Focus(ID int64) error {

	if ws, ok := w.list[w.focus]; ok {
		err := toggleButtonFocus(ws.Button, false)
		if err != nil {
			return err
		}
	}

	if wsb, ok := w.list[ID]; ok {
		w.focus = ID
		err := toggleButtonFocus(wsb.Button, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WorkspaceBox) FocusOrAdd(name string, ID int64) error {
	if w.focus == ID {
		return nil
	}

	if ws, ok := w.list[w.focus]; ok {
		toggleButtonFocus(ws.Button, false)
	}

	if ws, ok := w.list[ID]; ok {
		w.focus = ID
		toggleButtonFocus(ws.Button, true)
		return nil
	}

	return w.Add(name, ID)

}

func toggleButtonFocus(btn *gtk.Button, focus bool) error {
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
