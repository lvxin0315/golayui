package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const (
	CheckBox = "checkbox"
	Switch   = "switch"
)

type CheckboxOptionWeight struct {
	Title string
	Attr
	IsOld bool
}

type CheckboxWeight struct {
	OptionList []*CheckboxOptionWeight
}

func (s *CheckboxWeight) Output() (string, error) {
	return TemplateParse(s)
}

func (s *CheckboxWeight) GetTpl() string {
	return tpl.CheckboxTpl
}

func (s *CheckboxWeight) GetFormItemWeightType() string {
	return CheckBox
}

type SwitchWeight struct {
	Attr
}

func (s *SwitchWeight) Output() (string, error) {
	return TemplateParse(s)
}

func (s *SwitchWeight) GetTpl() string {
	return tpl.SwitchTpl
}

func (s *SwitchWeight) GetFormItemWeightType() string {
	return Switch
}
