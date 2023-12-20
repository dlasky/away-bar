package internal

import (
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type IconType string

const (
	StaticIcon  IconType = "static"
	DynamicIcon IconType = "dynamic"
)

type DynamicIconMap = map[string][]int

type Icon struct {
	iconType IconType
	icon     *gtk.Image
	config   *DynamicIconMap
	templ    *Tmpl
}

type IconConfig struct {
	IconType      IconType        `hcl:"type"`
	Path          *string         `hcl:"path,optional"`
	Config        *DynamicIconMap `hcl:"config,optional"`
	ValueTemplate *string         `hcl:"template,optional"`
}

func NewIconFromConfig(cfg *IconConfig) (*Icon, error) {
	spew.Dump(cfg)
	if cfg.IconType == StaticIcon {
		return NewStaticIcon(*cfg.Path)
	}
	return NewDynamicIcon(*cfg.Config, *cfg.ValueTemplate)
}

func (i *Icon) GetWidget() *gtk.Image {
	return i.icon
}

func getIconBase(iconType IconType, path string) (*Icon, error) {
	icon, err := gtk.ImageNewFromFile(path)
	if err != nil {
		return nil, err
	}

	return &Icon{
		iconType,
		icon,
		nil,
		nil,
	}, nil
}

func NewStaticIcon(path string) (*Icon, error) {
	return getIconBase(StaticIcon, path)
}

func NewDynamicIcon(config DynamicIconMap, templateRaw string) (*Icon, error) {
	icon, err := getIconBase(DynamicIcon, "")
	if err != nil {
		return nil, err
	}
	icon.config = &config
	tmp, err := NewTmpl("i", templateRaw)
	if err != nil {
		return nil, err
	}
	icon.templ = tmp
	return icon, err
}

func (i *Icon) Render(value any) error {
	//TODO: replace with an interval tree
	var candidate string

	preValue, err := i.templ.Render(value)
	if err != nil {
		return err
	}
	val, err := strconv.ParseInt(preValue, 10, 32)
	if err != nil {
		return err
	}

	for path, interval := range *i.config {
		if int(val) > interval[0] && int(val) < interval[1] {
			candidate = path
		}
	}
	glib.IdleAdd(func() {
		i.icon.SetFromFile(candidate)
	})
	return nil
}
