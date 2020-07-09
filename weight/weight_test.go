package weight

import (
	"io/ioutil"
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
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children,
		//按钮
		&ButtonWeight{
			Title:  "测试按钮",
			Style:  BtnDanger,
			Size:   BtnLg,
			Radius: true,
		},
		&IconButtonWeight{
			Icon:  IconIos,
			Style: BtnPrimary,
			Size:  BtnXs,
		},
		&FluidButtonWeight{
			Title: "测试流按钮",
			Style: "",
			Size:  "",
		},
	)
	//添加按钮组
	btnGroup := new(BtnGroupWeight)
	btnGroup.Children = append(btnGroup.Children, &ButtonWeight{
		Title: "测试1",
	}, &ButtonWeight{
		Title: "测试2",
	}, &ButtonWeight{
		Title: "测试3",
	})
	//添加按钮容器
	btnContainer := new(BtnContainerWeight)
	btnContainer.Children = append(btnContainer.Children, &ButtonWeight{
		Title: "测试4",
	}, &ButtonWeight{
		Title: "测试5",
	}, &ButtonWeight{
		Title: "测试6",
	})
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, btnGroup, btnContainer)

	//添加表单
	form := new(FormWeight)
	//基础的input
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, form)
	form.Children = append(form.Children, &FormItemWeight{
		Label: "单行输入框",
		Item: &InputTextWeight{
			Attr: Attr{
				Name: "danhangName",
				Id:   "danhangId",
			},
			Placeholder: "请输入标题",
		},
	}, &FormItemWeight{
		Label: "单行密码框",
		Item: &InputPasswordWeight{
			Attr: Attr{
				Name: "pn1",
				Id:   "pi1",
			},
			Placeholder: "请输入密码",
		},
	}, &FormItemWeight{
		Label: "单行日期",
		Item: &InputTextWeight{
			Attr: Attr{
				Name: "rn1",
				Id:   "ri1",
			},
			Placeholder: "yyyy-MM-dd",
			IsDate:      true,
		},
	}, &FormInlineItemWeight{
		InlineFormItemWeightList: []*FormItemWeight{
			{
				Label: "验证手机",
				Item: &InputTextWeight{
					Attr: Attr{
						Name: "n1",
						Id:   "i1",
					},
					Placeholder: "请输入手机",
				},
			},
			{
				Label: "验证邮箱",
				Item: &InputTextWeight{
					Attr: Attr{
						Name: "n2",
						Id:   "i2",
					},
					Placeholder: "请输入邮箱",
				},
			},
		},
	})
	//select
	form.Children = append(form.Children, &FormItemWeight{
		Label: "单行选择",
		Item: &SelectWeight{
			Attr: Attr{
				Name: "sn1",
				Id:   "si1",
			},
			OptionList: []*SelectOptionWeight{
				{
					Title: "t1",
					Value: "v1",
				},
				{
					Title: "t2",
					Value: "v2",
				},
				{
					Title: "t3",
					Value: "v3",
				},
				{
					Title: "t4",
					Value: "v4",
				},
				{
					Title: "t5",
					Value: "v5",
				},
				{
					Title: "t6",
					Value: "v6",
				},
			},
		},
	})

	//checkbox
	form.Children = append(form.Children, &FormItemWeight{
		Label: "单行选择",
		Item: &CheckboxWeight{
			OptionList: []*CheckboxOptionWeight{
				{
					Title: "c1",
					Name:  "c1",
				},
				{
					Title: "c2",
					Name:  "c2",
				},
				{
					Title: "c3",
					Name:  "c3",
				},
				{
					Title: "c4",
					Name:  "c4",
				},
				{
					Title: "c5",
					Name:  "c5",
				},
				{
					Title: "c6",
					Name:  "c6",
				},
			},
		},
	})

	//checkbox
	form.Children = append(form.Children, &FormItemWeight{
		Label: "单行选择(旧)",
		Item: &CheckboxWeight{
			OptionList: []*CheckboxOptionWeight{
				{
					Title: "o1",
					Name:  "o1",
					IsOld: true,
				},
				{
					Title: "o2",
					Name:  "o2",
					IsOld: true,
				},
				{
					Title: "o3",
					Name:  "o3",
					IsOld: true,
				},
				{
					Title: "o4",
					Name:  "o4",
					IsOld: true,
				},
				{
					Title: "o5",
					Name:  "o5",
					IsOld: true,
				},
				{
					Title: "o6",
					Name:  "o6",
					IsOld: true,
				},
			},
		},
	})

	//radio
	form.Children = append(form.Children, &FormItemWeight{
		Label: "单选框",
		Item: &RadioWeight{
			OptionList: []*RadioOptionWeight{
				{
					Title: "c1",
					Name:  "c1",
				},
				{
					Title: "c2",
					Name:  "c2",
				},
				{
					Title: "c3",
					Name:  "c3",
				},
				{
					Title: "c4",
					Name:  "c4",
				},
				{
					Title: "c5",
					Name:  "c5",
				},
				{
					Title: "c6",
					Name:  "c6",
				},
			},
		},
	})

	//textarea
	form.Children = append(form.Children, &FormItemWeight{
		Label: "文本域",
		Item: &TextareaWeight{
			Attr: Attr{
				Name: "tan1",
				Id:   "tai1",
			},
			Placeholder: "tap1",
		},
	})

	//tab
	tabItem1 := new(TabItemWeight)
	tabItem2 := new(TabItemWeight)
	tabItem3 := new(TabItemWeight)
	tabItem1.Title = "T1"
	tabItem2.Title = "T2"
	tabItem3.Title = "T3"
	tabItem1.Children = append(tabItem1.Children, &TextWeight{
		Text: "tab1-1 content",
	}, &TextWeight{
		Text: "tab1-2 content",
	})
	tabItem2.Children = append(tabItem2.Children, btnGroup)
	tabItem3.Children = append(tabItem3.Children, btnContainer)
	//普通tab
	tabWeight := new(TabWeight)
	tabWeight.ItemList = append(tabWeight.ItemList, tabItem1, tabItem2, tabItem3)
	//带样式tab
	tab1Weight := new(TabWeight)
	tab1Weight.ItemList = append(tab1Weight.ItemList, tabItem1, tabItem2, tabItem3)
	tab1Weight.Style = TabCard
	//可以关闭tab
	tab2Weight := new(TabWeight)
	tab2Weight.ItemList = append(tab2Weight.ItemList, tabItem1, tabItem2, tabItem3)
	tab2Weight.Style = TabBrief
	tab2Weight.AllowClose = true
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, tabWeight, tab1Weight, tab2Weight)

	//table
	tableWeight := new(TableWeight)
	tableWeight.Size = TableSizeSm
	tableWeight.Style = TableStyleLine
	tableWeight.Colgroup = append(tableWeight.Colgroup, "", "100", "100", "200")
	tableWeight.Thead = &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "表头1"},
		{Content: "表头2"},
		{Content: "表头3"},
		{Content: "表头4"},
	}}
	tableWeight.Tbody = append(tableWeight.Tbody, &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "1内容1"},
		{Content: "1内容2"},
		{Content: "1内容3"},
		{Content: "1内容4"},
	}}, &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "2内容1"},
		{Content: "2内容2"},
		{Content: "2内容3"},
		{Content: "2内容4"},
	}}, &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "3内容1"},
		{Content: "3内容2"},
		{Content: "3内容3"},
		{Content: "3内容4"},
	}}, &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "4内容1"},
		{Content: "4内容2"},
		{Content: "4内容3"},
		{Content: "4内容4"},
	}}, &TableTrWeight{TdList: []*TableTdWeight{
		{Content: "5内容1"},
		{Content: "5内容2"},
		{Content: "5内容3"},
		{Content: "5内容4"},
	}})
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, tableWeight)
	//引用区块
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, &BlockquoteWeight{
		Content: "引用区域的文字",
	})
	//横线
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, &HrWeight{Style: HrStyleGreen})
	//字段集区块
	adminLayoutBodyWeight.Children = append(adminLayoutBodyWeight.Children, &FieldsetWeight{
		Title:   "引用区域的文字",
		Content: "内容区域。",
	})

	//Footer
	adminLayoutFooterWeight.Children = append(adminLayoutFooterWeight.Children, &TextWeight{Text: " © layui.com - 底部固定区域"})
	mainHtml, err := adminLayoutWeight.Output()
	if err != nil {
		t.Error(err)
	}
	ioutil.WriteFile("/Users/lvxin/go/src/github.com/lvxin0315/golayui/tmp/m.html", []byte(mainHtml), 0777)
}
