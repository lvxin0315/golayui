package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const (
	Radio = "radio"
)

type RadioOptionWeight struct {
	Title string
	Name  string
}

type RadioWeight struct {
	OptionList []*RadioOptionWeight
}

func (s *RadioWeight) Output() (string, error) {
	return TemplateParse(s)
}

func (s *RadioWeight) GetTpl() string {
	return tpl.RadioTpl
}

func (s *RadioWeight) GetFormItemWeightType() string {
	return Radio
}
