package weight

const (
	IconHeartFill = `layui-icon-heart-fill`
	IconIos       = `layui-icon-ios`
	//TODO 其他后续补充， https://www.layui.com/doc/element/icon.html
)

type TextWeight struct {
	Text string
}

func (t *TextWeight) Output() (string, error) {
	return t.Text, nil
}

func (t *TextWeight) GetTpl() string {
	return ""
}
