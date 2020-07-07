package weight

import "github.com/lvxin0315/golayui/weight/tpl"

//按钮样式
const (
	BtnPrimary  = `layui-btn-primary`
	BtnNormal   = `layui-btn-normal`
	BtnWarm     = `layui-btn-warm`
	BtnDanger   = `layui-btn-danger`
	BtnDisabled = `layui-btn-disabled`
)

//按钮尺寸
const (
	BtnLg = `layui-btn-lg`
	BtnSm = `layui-btn-sm`
	BtnXs = `layui-btn-xs`
)

type Button struct {
	Title  string
	Style  string
	Size   string
	Radius bool
}

func (b *Button) Output() (string, error) {
	return TemplateParse(b)
}

func (b *Button) GetTpl() string {
	return tpl.ButtonTpl
}

//流体按钮
type FluidButton struct {
	Title  string
	Style  string
	Size   string
	Radius bool
}

func (b *FluidButton) Output() (string, error) {
	return TemplateParse(b)
}

func (b *FluidButton) GetTpl() string {
	return tpl.FluidButtonTpl
}

//流体按钮
type IconButton struct {
	Icon   string
	Style  string
	Size   string
	Radius bool
}

func (b *IconButton) Output() (string, error) {
	return TemplateParse(b)
}

func (b *IconButton) GetTpl() string {
	return tpl.IconButtonTpl
}
