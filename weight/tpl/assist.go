package tpl

const BlockquoteTpl = `<blockquote class="layui-elem-quote {{.Style}}">{{.Content}}</blockquote>`

const FieldsetTpl = `<fieldset class="layui-elem-field {{.Style}}">
  <legend>{{.Title}}</legend>
  <div class="layui-field-box">
    {{.Content}}
  </div>
</fieldset>`

const HrTpl = `<hr class="{{.Style}}">`
