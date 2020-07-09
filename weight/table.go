package weight

import "github.com/lvxin0315/golayui/weight/tpl"

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
