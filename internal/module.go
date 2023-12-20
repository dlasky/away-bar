package internal

import (
	"fmt"
	"log"

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
	icon            *Icon
	template        *Tmpl
	tooltipTemplate *Tmpl
}

type ModuleConfig struct {
	Name    string      `hcl:"title"`
	Format  string      `hcl:"format"`
	Tooltip string      `hcl:"tooltip"`
	Icon    *IconConfig `hcl:"icon,block"`
}

func NewModuleFromConfig(cfg *ModuleConfig) (*Module, error) {
	var icn *Icon

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}

	ctx, err := box.GetStyleContext()
	if err != nil {
		return nil, err
	}
	ctx.AddClass("module")
	ctx.AddClass(cfg.Name)

	label, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}
	label.SetName(cfg.Name)

	if cfg.Icon != nil {
		icn, err := NewIconFromConfig(cfg.Icon)
		if err != nil {
			return nil, err
		}
		box.Add(icn.GetWidget())
	}

	box.Add(label)
	box.ShowAll()

	tmp, err := NewTmpl(cfg.Name, cfg.Format)
	if err != nil {
		fmt.Println("compile error")
		return nil, err
	}

	tmp2, err := NewTmpl(cfg.Name+"tt", cfg.Tooltip)
	if err != nil {
		return nil, err
	}

	return &Module{
		name:            cfg.Name,
		label:           label,
		icon:            icn,
		box:             box,
		template:        tmp,
		tooltipTemplate: tmp2,
	}, nil
}

func (l *Module) GetWidget() gtk.IWidget {
	return l.box
}

func (l *Module) Render(value interface{}) error {
	text, err := l.template.Render(value)
	if err != nil {
		fmt.Println("render error")
		return err
	}
	tooltip, err := l.tooltipTemplate.Render(value)
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

func (l *Module) Error(err error) {
	fmt.Printf("[%v]: %v\n", l.name, err)
}

func (l *Module) Fatal(err error) {
	log.Fatalf("[%v]: %v\n", l.name, err)
}
