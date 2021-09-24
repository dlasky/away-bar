package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type IModule interface {
	GetWidget() gtk.IWidget
	Render(value interface{}) error
}

type Module struct {
	name            string
	label           *gtk.Label
	box             *gtk.Box
	icon            *gtk.Image
	templateRaw     string
	template        *template.Template
	tooltipRaw      string
	tooltipTemplate *template.Template
}

func NewModule(name string, templateRaw string, tooltipTemplate string, iconPath string) (*Module, error) {

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}

	ctx, err := box.GetStyleContext()
	if err != nil {
		return nil, err
	}
	ctx.AddClass("module")
	ctx.AddClass(name)

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
		fmt.Println("compile error")
		return nil, err
	}

	t2 := template.New(name)
	tmp2, err := t2.Parse(tooltipTemplate)
	if err != nil {
		return nil, err
	}

	return &Module{
		name:            name,
		label:           label,
		icon:            icon,
		box:             box,
		templateRaw:     templateRaw,
		template:        tmp,
		tooltipRaw:      tooltipTemplate,
		tooltipTemplate: tmp2,
	}, nil
}

func (l *Module) GetWidget() gtk.IWidget {
	return l.box
}

func (l *Module) Render(value interface{}) error {
	text, err := render(*l.template, value)
	if err != nil {
		fmt.Println("render error")
		return err
	}
	tooltip, err := render(*l.tooltipTemplate, value)
	if err != nil {
		return err
	}

	glib.IdleAdd(func() bool {
		l.label.SetText(text)
		l.box.SetTooltipText(tooltip)
		return false
	})
	return err
}

func render(tmp template.Template, value interface{}) (string, error) {
	b := bytes.NewBuffer([]byte{})
	err := tmp.Execute(b, value)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (l *Module) error(err error) {
	fmt.Printf("[%v]: %v\n", l.name, err)
}

func (l *Module) fatal(err error) {
	log.Fatalf("[%v]: %v\n", l.name, err)
}
