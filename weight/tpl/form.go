package tpl

const FormTpl = `<form class="layui-form" action="{{.Action}}" lay-filter="{{.Id}}">{{.ChildrenHtml}}</form>`

const FormItemTpl = `<div class="layui-form-item">
<label class="layui-form-label">{{.Label}}</label>
	<div class="layui-input-block">
	  {{.ItemHtml}}
	</div>
</div>`

const FromInlineItemTpl = `<div class="layui-form-item">
{{range .InlineFormItemWeightList}}
<div class="layui-inline">
  <label class="layui-form-label">{{.Label}}</label>
  <div class="layui-input-inline">
	{{.ItemHtml}}
  </div>
</div>
{{end}}
</div>`

const FormLabelTpl = `<div class="layui-form-item">
<label class="layui-form-label">{{.Label}}</label>
<div class="layui-input-block">
{{.FormHtml}}
</div>
</div>`
