package weight

import (
	"github.com/lvxin0315/golayui/weight/tpl"
)

const (
	Select = "select"
)

type SelectOptionWeight struct {
	Title string
	Value string
}

type SelectWeight struct {
	Attr
	OptionList []*SelectOptionWeight
}

func (s *SelectWeight) Output() (string, error) {
	return TemplateParse(s)
}

func (s *SelectWeight) GetTpl() string {
	return tpl.SelectTpl
}

func (s *SelectWeight) GetFormItemWeightType() string {
	return Select
}
