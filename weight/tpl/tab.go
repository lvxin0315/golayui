package tpl

const TabTpl = `<div class="layui-tab {{.Style}}" {{if .AllowClose}} lay-allowClose="true" {{end}}>
  <ul class="layui-tab-title">
	{{range .ItemList}}
		<li>{{.Title}}</li>
	{{end}}
  </ul>
  <div class="layui-tab-content">
	{{range .ItemList}}
		<div class="layui-tab-item">{{.ChildrenHtml}}</div>
	{{end}}
  </div>
</div>`

const TabItem = `{{.ChildrenHtml}}`
