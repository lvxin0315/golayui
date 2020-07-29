package tpl

const UploadTpl = `<button class="layui-btn" id="{{.Id}}" lay-data="{url: '/', acceptMime: 'image/*', multiple: true, done: function(res, index, upload){} }">{{.Title}}</button>`
