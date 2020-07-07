package tpl

const ButtonTpl = `<button type="button" class="layui-btn {{if .Radius}} layui-btn-radius {{end}} {{.Style}} {{.Size}}">{{.Title}}</button>`
const FluidButtonTpl = `<button type="button" class="layui-btn {{if .Radius}} layui-btn-radius {{end}} layui-btn-fluid {{.Style}} {{.Size}}">{{.Title}}</button>`
const IconButtonTpl = `<button type="button" class="layui-btn {{if .Radius}} layui-btn-radius {{end}} {{.Style}} {{.Size}}"><i class="layui-icon {{.Icon}}"></i></button>`
const BtnGroupTpl = `<div class="layui-btn-group">{{.ChildrenHtml}}</div>`
const BtnContainerTpl = `<div class="layui-btn-container">{{.ChildrenHtml}}</div>`
