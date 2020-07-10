package weight

import "github.com/lvxin0315/golayui/weight/tpl"

//面板
type CardWeight struct {
	Title string
	FullStateWeight
}

func (c *CardWeight) Output() (string, error) {
	return c.FullStateWeight.TemplateParse(c)
}

func (c *CardWeight) GetTpl() string {
	return tpl.CardTpl
}

//折叠面板
type CollaWeight struct {
	Accordion     bool
	CollaItemList []*CardWeight
}

func (c *CollaWeight) Output() (string, error) {
	//子集渲染
	for _, item := range c.CollaItemList {
		_, err := item.Output()
		if err != nil {
			return "", err
		}
	}
	return TemplateParse(c)
}

func (c *CollaWeight) GetTpl() string {
	return tpl.CollaTpl
}
