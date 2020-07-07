package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const (
	InputText     = "text"
	InputPassword = "password"
)

type FormItemWeightImpl interface {
	Weight
	SetName(name string)
	SetId(id string)
	GetFormItemWeightType() string
}

type InputTextWeight struct {
	Attr
	Placeholder string
	IsDate      bool
}

func (i *InputTextWeight) Output() (string, error) {
	return TemplateParse(i)
}

func (i *InputTextWeight) GetTpl() string {
	return tpl.InputTextTpl
}

func (i *InputTextWeight) SetName(name string) {
	i.Name = name
}

func (i *InputTextWeight) SetId(id string) {
	i.Id = id
}

func (i *InputTextWeight) GetFormItemWeightType() string {
	return InputText
}

type InputPasswordWeight struct {
	Attr
	Placeholder string
}

func (i *InputPasswordWeight) Output() (string, error) {
	return TemplateParse(i)
}

func (i *InputPasswordWeight) GetTpl() string {
	return tpl.InputPasswordTpl
}

func (i *InputPasswordWeight) SetName(name string) {
	i.Name = name
}

func (i *InputPasswordWeight) SetId(id string) {
	i.Id = id
}

func (i *InputPasswordWeight) GetFormItemWeightType() string {
	return InputPassword
}
