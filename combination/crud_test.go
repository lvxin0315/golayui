package combination

import (
	"github.com/lvxin0315/golayui/weight"
	"testing"
)

var savePath = `/Users/lvxin/go/src/github.com/lvxin0315/golayui/tmp/`

func TestBuild(t *testing.T) {
	var filedItemList []*Item
	filedItemList = append(filedItemList, &Item{
		Title: "名称",
		Field: "name",
		Type:  weight.InputText,
	}, &Item{
		Title: "密码",
		Field: "password",
		Type:  weight.InputPassword,
	}, &Item{
		Title: "性别",
		Field: "sex",
		Type:  weight.Select,
	}, &Item{
		Title: "签名",
		Field: "sign",
		Type:  weight.Textarea,
	}, &Item{
		Title: "字段1",
		Field: "f1",
		Type:  weight.CheckBox,
	}, &Item{
		Title: "字段2",
		Field: "f2",
		Type:  weight.Radio,
	})
	Build(filedItemList, savePath)
}
