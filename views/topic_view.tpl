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
          body {margin-top: 50px;}
            <h1>{{.Topics.Title}}<a href="/topic/modify?tid={{.Tid}}" class="btn btn-default">修改文章</a></h1>
            {{.Topics.Content}}
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
