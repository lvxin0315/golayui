package tpl

const TextareaTpl = `<textarea placeholder="{{.Placeholder}}" name="{{.Name}}" id="{{.Id}}" class="layui-textarea">{{if .Content}}{{.Content}}{{end}}</textarea>`
