package weight

import (
	"fmt"
	"github.com/lvxin0315/golayui/weight/tpl"
)

const (
	ColXS          = "xs"
	ColSM          = "sm"
	ColMD          = "md"
	ColLG          = "lg"
	ColStyle       = `layui-col-%s%d `
	ColOffsetStyle = `layui-col-%s-offset%d `
)

type ContainerWeight struct {
	IsFluid bool
	FullStateWeight
}

func (c *ContainerWeight) Output() (string, error) {
	return c.FullStateWeight.TemplateParse(c)
}

func (c *ContainerWeight) GetTpl() string {
	return tpl.ContainerTpl
}

type RowWeight struct {
	FullStateWeight
}

func (r *RowWeight) Output() (string, error) {
	return r.FullStateWeight.TemplateParse(r)
}

func (r *RowWeight) GetTpl() string {
	return tpl.RowTpl
}

type ColWeight struct {
	FullStateWeight
	XS        *ColSizeItem
	SM        *ColSizeItem
	MD        *ColSizeItem
	LG        *ColSizeItem
	StyleHtml string
}

func (c *ColWeight) setStyleHtml() {
	if c.LG != nil && c.LG.Num > 0 {
		c.StyleHtml += fmt.Sprintf(ColStyle, ColLG, c.LG.Num)
		if c.LG.OffsetNum > 0 {
			c.StyleHtml += fmt.Sprintf(ColOffsetStyle, ColLG, c.LG.OffsetNum)
		}
	}
	if c.MD != nil && c.MD.Num > 0 {
		c.StyleHtml += fmt.Sprintf(ColStyle, ColMD, c.MD.Num)
		if c.MD.OffsetNum > 0 {
			c.StyleHtml += fmt.Sprintf(ColOffsetStyle, ColMD, c.MD.OffsetNum)
		}
	}
	if c.SM != nil && c.SM.Num > 0 {
		c.StyleHtml += fmt.Sprintf(ColStyle, ColSM, c.SM.Num)
		if c.SM.OffsetNum > 0 {
			c.StyleHtml += fmt.Sprintf(ColOffsetStyle, ColSM, c.SM.OffsetNum)
		}
	}
	if c.XS != nil && c.XS.Num > 0 {
		c.StyleHtml += fmt.Sprintf(ColStyle, ColXS, c.XS.Num)
		if c.XS.OffsetNum > 0 {
			c.StyleHtml += fmt.Sprintf(ColOffsetStyle, ColXS, c.XS.OffsetNum)
		}
	}
}

func (c *ColWeight) Output() (string, error) {
	//先整理Style代码
	c.setStyleHtml()
	return c.FullStateWeight.TemplateParse(c)
}

func (c *ColWeight) GetTpl() string {
	return tpl.ColTpl
}

type ColSizeItem struct {
	Num       uint
	OffsetNum uint
}
