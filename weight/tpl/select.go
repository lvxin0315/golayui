package tpl

const SelectTpl = `<div class="layui-input-block">
  <select name="{{.Name}}" id="{{.Id}}" lay-filter="{{.Name}}">
    <option value=""></option>
	{{range .OptionList}}
	<option value="{{.Value}}">{{.Title}}</option>
	{{end}}
  </select>
</div>`
