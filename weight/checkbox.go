package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const (
	CheckBox = "checkbox"
)

type CheckboxOptionWeight struct {
	Title string
	Name  string
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
