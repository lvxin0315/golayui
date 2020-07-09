package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const Textarea = "textarea"

type TextareaWeight struct {
	Attr
	Placeholder string
}

func (t *TextareaWeight) Output() (string, error) {
	return TemplateParse(t)
}

func (t *TextareaWeight) GetTpl() string {
	return tpl.TextareaTpl
}

func (t *TextareaWeight) GetFormItemWeightType() string {
	return Textarea
}
