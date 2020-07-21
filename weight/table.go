package weight

import (
	"encoding/json"
	"github.com/lvxin0315/golayui/weight/tpl"
)

const (
	TableStyleLine = "line" //（行边框风格）
	TableStyleRow  = "row"  //（列边框风格）
	TableStyleNob  = "nob"  //（无边框风格）
	TableSizeSm    = "sm"   //（小尺寸）
	TableSizeLg    = "lg"   //（大尺寸）
)

type TableWeight struct {
	Style    string
	Size     string
	Colgroup []string
	Thead    *TableTrWeight
	Tbody    []*TableTrWeight
}

func (t *TableWeight) Output() (string, error) {
	return TemplateParse(t)
}

func (t *TableWeight) GetTpl() string {
	return tpl.TableTpl
}

type TableTrWeight struct {
	TdList []*TableTdWeight
}

type TableTdWeight struct {
	Content string
}

//数据table
type DataTableWeight struct {
	Attr
	Style               string
	Size                string
	FieldList           []*DataTableItem
	FieldListJsonString string
}

type DataTableItem struct {
	Field string `json:"field"`
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

func (t *DataTableWeight) Output() (string, error) {
	//toJsonString
	json, err := json.Marshal(t.FieldList)
	if err != nil {
		return "", err
	}
	t.FieldListJsonString = string(json)
	return TemplateParse(t)
}

func (t *DataTableWeight) GetTpl() string {
	return tpl.DataTableTpl
}
