package tpl

const CheckboxTpl = `{{range .OptionList}}
<input type="checkbox" name="{{.Name}}" title="{{.Title}}">
{{end}}`
