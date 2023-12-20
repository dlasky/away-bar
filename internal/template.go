package internal

import (
	"bytes"
	"fmt"
	"text/template"
)

type Tmpl struct {
	tmp *template.Template
}

func NewTmpl(name string, tmpl string) (*Tmpl, error) {

	t := template.New(name)
	tmp, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println("template error", name, err)
		return nil, err
	}
	return &Tmpl{
		tmp,
	}, nil
}

func (t *Tmpl) Render(value any) (string, error) {
	b := bytes.NewBuffer([]byte{})
	err := t.tmp.Execute(b, value)

	if err != nil {
		return "", err
	}

	return b.String(), nil
}
