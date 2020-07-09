package weight

import (
	"github.com/lvxin0315/golayui/weight/tpl"
)

const (
	TabBrief = "layui-tab-brief"
	TabCard  = "layui-tab-card"
)

//选项卡
type TabWeight struct {
	Style      string
	AllowClose bool
	ItemList   []*TabItemWeight
}

func (t *TabWeight) Output() (string, error) {
	//先把子集内容output
	for _, item := range t.ItemList {
		_, err := item.Output()
		if err != nil {
			return "", err
		}
	}
	return TemplateParse(t)
}

func (t *TabWeight) GetTpl() string {
	return tpl.TabTpl
}

type TabItemWeight struct {
	Title string
	FullStateWeight
}

func (t *TabItemWeight) Output() (string, error) {
	return t.FullStateWeight.TemplateParse(t)
}

func (t *TabItemWeight) GetTpl() string {
	return `{{.ChildrenHtml}}`
}
