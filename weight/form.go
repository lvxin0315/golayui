package weight

import (
	"github.com/lvxin0315/golayui/weight/tpl"
)

//表单
type FormWeight struct {
	Attr
	Action string
	FullStateWeight
}

func (f *FormWeight) Output() (string, error) {
	return f.FullStateWeight.TemplateParse(f)
}

func (f *FormWeight) GetTpl() string {
	return tpl.FormTpl
}

//表单元素,一行一个
type FormItemWeight struct {
	Label    string
	Item     FormItemWeightImpl
	ItemHtml string
}

func (f *FormItemWeight) Output() (string, error) {
	//先把元素渲染
	if f.Item != nil {
		itemHtml, err := f.Item.Output()
		if err != nil {
			return "", err
		}
		f.ItemHtml = itemHtml
	}
	return TemplateParse(f)
}

func (f *FormItemWeight) GetTpl() string {
	return tpl.FormItemTpl
}

//表单元素,一行多个
type FormInlineItemWeight struct {
	InlineFormItemWeightList []*FormItemWeight //虽然用单行结构，但是不用对应的output
}

func (f *FormInlineItemWeight) Output() (string, error) {
	//把子集的ItemHtml生成出来
	for _, item := range f.InlineFormItemWeightList {
		_, _ = item.Output()
	}
	return TemplateParse(f)
}

func (f *FormInlineItemWeight) GetTpl() string {
	return tpl.FromInlineItemTpl
}
