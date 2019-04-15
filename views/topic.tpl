{{template "header" .}}
        <title>文章 - 我的beego博客</title>
    </head>

    <body>
        body { padding-top: 100px;}
        <div class="navbar navbar-default navbar-fixed-top">
            <div class="container">
                {{template "navbar" .}}
            </div>
        </div>

        <div class="container">
          body {margin-top: 50px;}
          <h1>文章列表</h1>
          <a href="/topic/add" class="btn btn-default">添加文章</a>
          <table class="table table-striped">
              <thead>
                  <tr>
                      <th>#</th>
                      <th>文章名称</th>
                      <th>最后更新</th>
                      <th>浏览</th>
                      <th>回复数</th>
                      <th>删除</th>
                  </tr>
              </thead>
              <tbody>
                  {{range .Topics}}
                  <tr>
                    <th>{{.Id}}</th>
                    <th>{{.Title}}</th>
                    <th>{{.Updated}}</th>
                    <th>{{.Views}}</th>
                    <th>{{.ReplyCount}}</th>
                    <th><a href="/topic/delete?tid={{.Id}}">删除</a></th>
                  </tr>
                  {{end}}
              </tbody>
          </table>
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
