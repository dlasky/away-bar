package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Module struct {
	name        string
	label       *gtk.Label
	box         *gtk.Box
	icon        *gtk.Image
	templateRaw string
	template    *template.Template
}

// type ModuleRender func(*gtk.Box, error)

// func RegisterModule(typ string, render ModuleRender) {

// }

func NewModule(name string, templateRaw string, iconPath string) (*Module, error) {

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}

	label, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	label.SetName(name)

	icon, err := gtk.ImageNewFromFile(iconPath)
	if err != nil {
		return nil, err
	}

	box.Add(icon)
	box.Add(label)
	box.ShowAll()

	t := template.New(name)
	tmp, err := t.Parse(templateRaw)
	if err != nil {
		return nil, err
	}

	return &Module{
		name:        name,
		label:       label,
		icon:        icon,
		box:         box,
		templateRaw: templateRaw,
		template:    tmp,
	}, nil
}

func (l *Module) GetWidget() gtk.IWidget {
	return l.box
}

func (l *Module) Render(value interface{}) error {
	b := bytes.NewBuffer([]byte{})
	err := l.template.Execute(b, value)
	if err != nil {
		return err
	}
	s := b.String()
	_, err = glib.IdleAdd(l.label.SetText, s)
	return err
}

func (l *Module) error(err error) {
	fmt.Printf("[%v]: %v\n", l.name, err)
}

func (l *Module) fatal(err error) {
	log.Fatalf("[%v]: %v\n", l.name, err)
}
