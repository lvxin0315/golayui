package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const (
	BlockquoteStyleNM  = `layui-quote-nm`
	FieldsetStyleTitle = `layui-field-title`
	HrStyleRed         = `layui-bg-red`
	HrStyleOrange      = `layui-bg-orange`
	HrStyleGreen       = `layui-bg-green`
	HrStyleBlue        = `layui-bg-blue`
	HrStyleBlack       = `layui-bg-black`
	HrStyleGray        = `layui-bg-gray`
)

//引用区块
type BlockquoteWeight struct {
	Style   string
	Content string
}

func (b *BlockquoteWeight) Output() (string, error) {
	return TemplateParse(b)
}

func (b *BlockquoteWeight) GetTpl() string {
	return tpl.BlockquoteTpl
}

//字段集区块
type FieldsetWeight struct {
	Style   string
	Title   string
	Content string
}

func (f *FieldsetWeight) Output() (string, error) {
	return TemplateParse(f)
}

func (f *FieldsetWeight) GetTpl() string {
	return tpl.FieldsetTpl
}

//横线
type HrWeight struct {
	Style string
}

func (h *HrWeight) Output() (string, error) {
	return TemplateParse(h)
}

func (h *HrWeight) GetTpl() string {
	return tpl.HrTpl
}
