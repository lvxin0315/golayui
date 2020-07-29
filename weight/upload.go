package weight

import "github.com/lvxin0315/golayui/weight/tpl"

const Upload = "upload"

type UploadWeight struct {
	Attr
	Title string
}

func (i *UploadWeight) Output() (string, error) {
	return TemplateParse(i)
}

func (i *UploadWeight) GetTpl() string {
	return tpl.UploadTpl
}

func (i *UploadWeight) GetFormItemWeightType() string {
	return Upload
}
