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

type FullStateWeight struct {
	Children     []Weight
	ChildrenHtml string
}

func (w *FullStateWeight) TemplateParse(fullStateWeight Weight) (string, error) {
	childrenHtml := ""
	//是否包含child
	if len(w.Children) > 0 {
		for _, child := range w.Children {
			childHtml, err := child.Output()
			if err != nil {
				return "", err
			}
			childrenHtml += childHtml
		}
	}
	w.ChildrenHtml = childrenHtml
	commonHtml, err := TemplateParse(fullStateWeight)
	if err != nil {
		return "", err
	}
	return commonHtml, nil
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
