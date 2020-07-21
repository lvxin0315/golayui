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

const DataTableTpl = `<table id="{{.Id}}" lay-filter="lay{{.Id}}" lay-skin="{{.Style}}" lay-size="{{.Size}}"></table>
<script type="text/html" id="tool{{.Id}}">
	<div class="layui-btn-container">
		<button class="layui-btn layui-btn-sm" lay-event="add">添加</button>
		<button class="layui-btn layui-btn-sm" lay-event="delete">删除</button>
		<button class="layui-btn layui-btn-sm" lay-event="update">编辑</button>
	</div>
</script>

<script>
	var layerIndex;
	layui.use(['table', 'layer','jquery'], function () {
		var table = layui.table;
		var layer = layui.layer;
		var $ = layui.jquery;
		table.render({
			toolbar: '#tool{{.Id}}',
			elem: '#{{.Id}}',
			url: 'url', //数据接口
			page: true, //开启分页
			cols: [{{.FieldListJsonString}}]
		});
		//监听事件
		table.on('toolbar(lay{{.Id}})', function(obj){
			switch(obj.event){
				case 'add':
					layer.msg('添加');
					break;
				case 'delete':
					layer.msg('删除');
					break;
				case 'update':
					layer.msg('编辑');
					break;
			}
		});
	});


</script>`
