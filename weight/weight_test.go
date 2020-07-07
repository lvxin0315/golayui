package weight

import (
	"fmt"
	"testing"
)

func TestAdminLayoutWeight_Output(t *testing.T) {
	adminLayoutWeight := new(AdminLayoutWeight)
	adminLayoutHeaderWeight := new(AdminLayoutHeaderWeight)
	adminLayoutSideWeight := new(AdminLayoutSideWeight)
	adminLayoutBodyWeight := new(AdminLayoutBodyWeight)
	adminLayoutFooterWeight := new(AdminLayoutFooterWeight)
	adminLayoutWeight.Title = "测试Title"
	adminLayoutWeight.Header = adminLayoutHeaderWeight
	adminLayoutWeight.Side = adminLayoutSideWeight
	adminLayoutWeight.Body = adminLayoutBodyWeight
	adminLayoutWeight.Footer = adminLayoutFooterWeight
	//导航logo标题
	adminLayoutHeaderWeight.Title = "导航标题"
	//导航居左
	adminLayoutHeaderWeight.LeftMenu = append(adminLayoutHeaderWeight.LeftMenu,
		&AdminLayoutHeaderMenuWeight{
			Title:    "left菜单1",
			Href:     DefaultHref,
			Children: nil,
		}, &AdminLayoutHeaderMenuWeight{
			Title:    "left菜单2",
			Href:     "/baidu",
			Children: nil,
		}, &AdminLayoutHeaderMenuWeight{
			Title: "left菜单3",
			Href:  DefaultHref,
			Children: []struct {
				Title string
				Href  string
			}{
				{"left菜单3-1", "/a"},
				{"left菜单3-2", "/b"},
				{"left菜单3-3", "/c"},
			},
		})
	//导航居右
	adminLayoutHeaderWeight.RightMenu = append(adminLayoutHeaderWeight.RightMenu, &AdminLayoutHeaderMenuWeight{
		Title: "贤心",
		Href:  DefaultHref,
		Children: []struct {
			Title string
			Href  string
		}{
			{"基本资料", "/a"},
			{"安全设置", "/b"},
		},
	}, &AdminLayoutHeaderMenuWeight{
		Title:    "退了",
		Href:     DefaultHref,
		Children: nil,
	})
	//左侧导航菜单
	adminLayoutSideWeight.Menu = append(adminLayoutWeight.Side.Menu, &AdminLayoutSideMenuWeight{
		Title: "所有商品",
		Href:  DefaultHref,
		Children: []struct {
			Title string
			Href  string
		}{
			{"列表一", "/a"},
			{"列表二", "/b"},
			{"列表三", "/c"},
		},
	})
	//主体
	adminLayoutBodyWeight.BodyHtml = "你好"
	//Footer
	adminLayoutFooterWeight.FooterHtml = " © layui.com - 底部固定区域"

	mainHtml, err := adminLayoutWeight.Output()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(mainHtml)
}
