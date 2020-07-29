package tpl

const CheckboxTpl = `{{range .OptionList}}
<input type="checkbox" name="{{.Name}}" title="{{.Title}}" id="{{.Id}}" lay-skin="{{if .IsOld}}primary{{end}}" >
{{end}}`

const SwitchTpl = `<input type="checkbox" name="{{.Name}}" id="{{.Id}}" lay-skin="switch" >`
