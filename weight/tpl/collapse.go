package tpl

const CardTpl = `<div class="layui-card">
  <div class="layui-card-header">{{.Title}}</div>
  <div class="layui-card-body">
    {{.ChildrenHtml}}
  </div>
</div>`

const CollaTpl = `<div class="layui-collapse" {{if .Accordion}} lay-accordion {{end}}>
  {{range .CollaItemList}}
  <div class="layui-colla-item">
    <h2 class="layui-colla-title">{{.Title}}</h2>
    <div class="layui-colla-content layui-show">
		{{.ChildrenHtml}}
    </div>
  </div>
  {{end}}
</div>`
