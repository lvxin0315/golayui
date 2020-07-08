package tpl

const RadioTpl = `{{range .OptionList}}
<input type="radio" name="{{.Name}}" title="{{.Title}}" >
{{end}}`
