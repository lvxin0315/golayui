package tpl

const InputTextTpl = `<input type="text" name="{{.Name}}" id="{{.Id}}" lay-verify="{{if .IsDate}}date{{else}}required{{end}}" autocomplete="off" placeholder="{{.Placeholder}}" class="layui-input">`
const InputPasswordTpl = `<input type="password" name="{{.Name}}" id="{{.Id}}" lay-verify="required" autocomplete="off" placeholder="{{.Placeholder}}" class="layui-input">`
