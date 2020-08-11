package tpl

const AdminLayoutTpl = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/layui/css/layui.css">
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
{{.BodyHtml}}
</div>
<script src="/layui/layui.all.js"></script>
<script>
    //JavaScript代码区域
    layui.use('element', function(){
        var element = layui.element;

    });
</script>
</body>
</html>`

const AdminLayoutHeaderTpl = `<div class="layui-header">
<div class="layui-logo">{{.Title}}</div>
<!-- 头部区域（可配合layui已有的水平导航） -->
{{if .LeftMenu}}
	<ul class="layui-nav layui-layout-left">
		{{range .LeftMenu}}
			<li class="layui-nav-item">
			{{if .Children}}
				<a href="javascript:;">{{.Title}}</a>
				<dl class="layui-nav-child">
				{{range .Children}}
					<dd><a href="{{.Href}}">{{.Title}}</a></dd>
				{{end}}
				</dl>
				{{else}}
				<a href="{{.Href}}">{{.Title}}</a>
			{{end}}
			</li>
		{{end}}
	</ul>
{{end}}
{{if .RightMenu}}
	<ul class="layui-nav layui-layout-right">
		{{range .RightMenu}}
			<li class="layui-nav-item">
			{{if .Children}}
				<a href="javascript:;">{{.Title}}</a>
				<dl class="layui-nav-child">
					{{range .Children}}
						<dd><a href="{{.Href}}">{{.Title}}</a></dd>
					{{end}}
				</dl>
			{{else}}
				<a href="{{.Href}}">{{.Title}}</a>
			{{end}}
			</li>
		{{end}}
	</ul>
{{end}}
</div>`

const AdminLayoutSideTpl = `<div class="layui-side layui-bg-black">
<div class="layui-side-scroll">
	<!-- 左侧导航区域（可配合layui已有的垂直导航） -->
	{{if .Menu}}
		<ul class="layui-nav layui-nav-tree">
			{{range .Menu}}
				<li class="layui-nav-item">
				{{if .Children}}
					<a href="javascript:;">{{.Title}}</a>
					<dl class="layui-nav-child">
						{{range .Children}}
							<dd><a href="{{.Href}}">{{.Title}}</a></dd>
						{{end}}
					</dl>
				{{else}}
					<a href="{{.Href}}">{{.Title}}</a>
				{{end}}
				</li>
			{{end}}
		</ul>
	{{end}}
</div>
</div>`

const AdminLayoutBodyTpl = `<div class="layui-body">
	<!-- 内容主体区域 -->
	{{.ChildrenHtml}}
</div>`

const AdminLayoutFooterTpl = `<div class="layui-footer">
	<!-- 底部固定区域 -->
	{{.ChildrenHtml}}
</div>`
