package tpl

const ContainerTpl = `<div class="{{if .IsFluid}} layui-fluid {{else}} layui-container {{end}}">
{{.ChildrenHtml}}
</div>`

const RowTpl = `<div class="layui-row">
{{.ChildrenHtml}}
</div>`

const ColTpl = `<div class="{{.StyleHtml}}">
{{.ChildrenHtml}}
</div>`
