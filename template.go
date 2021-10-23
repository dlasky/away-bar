// template helper functions
package main

import (
	"fmt"
	"html/template"
)

type DualFloat float64

func (f DualFloat) String() string {
	return fmt.Sprintf("%.2f", float64(f))
}

var templateFuncs = template.FuncMap{
	"Float": func() {},
}
