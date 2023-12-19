package icon

import "github.com/gotk3/gotk3/glib"

type DynamicIconMap = map[string][]int

func NewDynamicIcon(config DynamicIconMap) (*Icon, error) {
	icon, err := getIconBase(DynamicIcon, "")
	icon.config = &config
	return icon, err
}

func (i *Icon) Render(value int) {
	//TODO: replace with an interval tree
	var candidate string
	for path, interval := range *i.config {
		if value > interval[0] && value < interval[1] {
			candidate = path
		}
	}
	glib.IdleAdd(func() {
		i.icon.SetFromFile(candidate)
	})
}
