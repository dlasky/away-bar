package icon

import "github.com/gotk3/gotk3/gtk"

func NewStaticIcon(path string) (*Icon, error) {
	return getIconBase(StaticIcon, path)
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
	}, nil
}
