{{template "header" .}}
        <title>添加文章 - 我的beego博客</title>
    </head>

    <body>
        body { padding-top: 100px;}
        <div class="navbar navbar-default navbar-fixed-top">
            <div class="container">
                {{template "navbar" .}}
            </div>
        </div>

        <div class="container">
          {{$label := .Labels}}
          body {margin-top: 50px;}
            <h1>{{.Topics.Title}}<a href="/topic/modify?tid={{.Tid}}" class="btn btn-default">修改文章</a></h1>
            <h5>
              {{range $label}}
              <a href="/?label={{.}}">{{.}}</a>
              {{end}}
            </h5>
            {{.Topics.Content}}
        </div>

        <div class="container">
          {{$tid := .Topics.Id}}
          {{$IsLogin := .IsLogin}}
          {{range .Replies}}
          <h3>{{.Name}}<small>{{.Created}}</small> {{if $IsLogin}}<a href="/reply/delete?tid={{$tid}}&rid={{.Id}}">删除</a>{{end}}</h3>
          {{.Content}}
          {{end}}
          <h3>本文回复<h3>
            <form method="post" action="/reply/add">
              <input type="hidden" name="tid" value="{{.Tid}}">
              <div class="form-group">
                <label>显示昵称: </label>
                <input type="text" class="form-control" name="nickname">
              </div>

              <div class="form-group">
                <label>内容: </label>
                <textarea name="content" id="" cols="30" rows="10" class="form-control"></textarea>
              </div>
              <button class="btn btn-default">提交回复</button>
            </form>
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
