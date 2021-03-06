{{template "header"}}
        <title>首页 - 我的beego博客</title>
    </head>

    <body>
        body { padding-top: 100px;}
        <div class="navbar navbar-default navbar-fixed-top">
            <div class="container">
                {{template "navbar" .}}
            </div>
        </div>

        <div class="container">
          <div class="row">
          <div class="col-md-9">
            <div class="page-header">
                <h1>无耻之徒</h1>
                <h2>*****************************</h2>
                {{range .Topics}}
                <h1>{{.Title}}<h1>
                <h6 class="text-muted">文章发表于 {{.Created}}，共有 {{.Views}}次浏览，{{.ReplyCount}}个评论</h6>
                <p>
                    {{.Content}}
                </p>
            </div>
            {{end}}
          </div>

          <div class="col-md-3">
            body {margin-top: 50px;}
            <h3>文章分类</h3>
            <ul>
              {{range .Categories}}
              <li><a href="/?cate={{.Title}}">{{.Title}}</a></li>
              {{end}}
            </ul>
          </div>
        </div>
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <body>
</html>
