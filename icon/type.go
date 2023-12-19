package icon

import "github.com/gotk3/gotk3/gtk"

type IconType string

const (
	StaticIcon  IconType = "static"
	DynamicIcon IconType = "dynamic"
)

type IModule interface {
	GetWidget() gtk.Image
	Render(value interface{}) error
}

type Icon struct {
	iconType IconType
	icon     *gtk.Image
	config   *DynamicIconMap
}
