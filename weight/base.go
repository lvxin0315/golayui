package weight

type TextWeight struct {
	Text string
}

func (t *TextWeight) Output() (string, error) {
	return t.Text, nil
}

func (t *TextWeight) GetTpl() string {
	return ""
}

type EmptyWeight struct {
	FullStateWeight
}

func (w *EmptyWeight) Output() (string, error) {
	return w.FullStateWeight.TemplateParse(w)
}

func (w *EmptyWeight) GetTpl() string {
	return `{{.ChildrenHtml}}`
}
