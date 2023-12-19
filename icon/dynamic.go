package icon

type DynamicIconMap = map[int]string

func NewDynamicIcon(config DynamicIconMap) (*Icon, error) {
	icon, err := getIconBase(DynamicIcon, "")
	icon.config = &config
	return icon, err
}

func (i *Icon) Render(Value any) {

}
