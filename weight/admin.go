package weight

import (
	"github.com/lvxin0315/golayui/weight/tpl"
)

//管理端组件

//管理界面框架
type AdminLayoutWeight struct {
	Title    string //标题
	Header   *AdminLayoutHeaderWeight
	Side     *AdminLayoutSideWeight
	Body     *AdminLayoutBodyWeight
	Footer   *AdminLayoutFooterWeight
	BodyHtml string
}

func (w *AdminLayoutWeight) GetTpl() string {
	return tpl.AdminLayoutTpl
}

func (w *AdminLayoutWeight) Output() (string, error) {
	headerHtml, err := w.Header.Output()
	if err != nil {
		return "", err
	}
	sideHtml, err := w.Side.Output()
	if err != nil {
		return "", err
	}
	bodyHtml, err := w.Body.Output()
	if err != nil {
		return "", err
	}
	footerHtml, err := w.Footer.Output()
	if err != nil {
		return "", err
	}
	w.BodyHtml = headerHtml + sideHtml + bodyHtml + footerHtml
	return TemplateParse(w)
}

//顶部导航
type AdminLayoutHeaderWeight struct {
	LessStateWeight
	Title     string
	LeftMenu  []*AdminLayoutHeaderMenuWeight
	RightMenu []*AdminLayoutHeaderMenuWeight
}

func (w *AdminLayoutHeaderWeight) GetTpl() string {
	return tpl.AdminLayoutHeaderTpl
}

func (w *AdminLayoutHeaderWeight) Output() (string, error) {
	return TemplateParse(w)
}

//顶部导航菜单
type AdminLayoutHeaderMenuWeight struct {
	Title    string
	Href     string
	Children []struct {
		Title string
		Href  string
	}
}

//左侧导航区域
type AdminLayoutSideWeight struct {
	LessStateWeight
	Menu []*AdminLayoutSideMenuWeight
}

func (w *AdminLayoutSideWeight) GetTpl() string {
	return tpl.AdminLayoutSideTpl
}

func (w *AdminLayoutSideWeight) Output() (string, error) {
	return TemplateParse(w)
}

type AdminLayoutSideMenuWeight struct {
	Title    string
	Href     string
	Children []struct {
		Title string
		Href  string
	}
}

//内容主体区域
type AdminLayoutBodyWeight struct {
	FullStateWeight
}

func (w *AdminLayoutBodyWeight) GetTpl() string {
	return tpl.AdminLayoutBodyTpl
}

func (w *AdminLayoutBodyWeight) Output() (string, error) {
	return w.FullStateWeight.TemplateParse(w)
}

//底部固定区域
type AdminLayoutFooterWeight struct {
	FullStateWeight
}

func (w *AdminLayoutFooterWeight) GetTpl() string {
	return tpl.AdminLayoutFooterTpl
}

func (w *AdminLayoutFooterWeight) Output() (string, error) {
	return w.FullStateWeight.TemplateParse(w)
}
