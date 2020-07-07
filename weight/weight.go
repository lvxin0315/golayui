package weight

import (
	"strings"
	"text/template"
)

const (
	DefaultHref = "javascript:;"
)

//组件基础
type Weight interface {
	Output() (string, error)
	GetTpl() string
}

type LessStateWeight struct {
}

type FullStateWeight struct {
	Children []Weight
}

func TemplateParse(w Weight) (string, error) {
	tmpl, err := template.New("").Parse(w.GetTpl())
	if err != nil {
		return "", err
	}
	strBuilder := new(strings.Builder)
	err = tmpl.Execute(strBuilder, w)
	if err != nil {
		return "", err
	}
	return strBuilder.String(), nil
}
