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

type ButtonWeight struct {
	Attr
	Title  string
	Style  string
	Size   string
	Radius bool
}

func (b *ButtonWeight) Output() (string, error) {
	return TemplateParse(b)
}

func (b *ButtonWeight) GetTpl() string {
	return tpl.ButtonTpl
}

//流体按钮
type FluidButtonWeight struct {
	Attr
	Title  string
	Style  string
	Size   string
	Radius bool
}

func (b *FluidButtonWeight) Output() (string, error) {
	return TemplateParse(b)
}

func (b *FluidButtonWeight) GetTpl() string {
	return tpl.FluidButtonTpl
}

//流体按钮
type IconButtonWeight struct {
	Attr
	Icon   string
	Style  string
	Size   string
	Radius bool
}

func (b *IconButtonWeight) Output() (string, error) {
	return TemplateParse(b)
}

func (b *IconButtonWeight) GetTpl() string {
	return tpl.IconButtonTpl
}

//按钮容器
type BtnContainerWeight struct {
	FullStateWeight
}

func (w *BtnContainerWeight) GetTpl() string {
	return tpl.BtnContainerTpl
}

func (w *BtnContainerWeight) Output() (string, error) {
	return w.FullStateWeight.TemplateParse(w)
}

//按钮组
type BtnGroupWeight struct {
	FullStateWeight
}

func (w *BtnGroupWeight) GetTpl() string {
	return tpl.BtnGroupTpl
}

func (w *BtnGroupWeight) Output() (string, error) {
	return w.FullStateWeight.TemplateParse(w)
}
