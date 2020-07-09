package tpl

const TableTpl = `<table class="layui-table" lay-skin="{{.Style}}" lay-size="{{.Size}}">
  {{if .Colgroup}}
  <colgroup>
    {{range $cw := .Colgroup}}
	<col width="{{$cw}}">
	{{end}}
  </colgroup>
  {{end}}

  {{if .Thead}}
  <thead>
    <tr>
      {{range .Thead.TdList}}
		<th>{{.Content}}</th>
      {{end}}
    </tr> 
  </thead>
  {{end}}
  {{if .Tbody}}
  <tbody>
	{{range .Tbody}}
	<tr>
      {{range .TdList}}
		<td>{{.Content}}</td>
      {{end}}
    </tr>
	{{end}}
  </tbody>
  {{end}}
</table>`
