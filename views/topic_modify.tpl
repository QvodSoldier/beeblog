{{template "header" .}}
        <title>修改文章 - 我的beego博客</title>
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
          <h1>修改文章</h1>
          <form method="post" action="/topic">
              <input type="hidden" name="tid" value="{{.Tid}}">
              <div class="form-group">
                <label>文章标题: </label>
                <input type="text" name="title" class="form-control" value="{{.Topic.Title}}">
              </div>

              <div class="form-group">
                <label>文章内容: </label>
                <textarea name="content" clos="30" rows="10" class="form-control">{{.Topic.Content}}</textarea>
              </div>

              <div class="form-group">
                <label>文章标签: </label>
                <input type="text" name="label" class="form-control" value="{{.Topic.Labels}}">
              </div>

              <div class="form-group">
                <label>文章分类: </label>
                <input type="text" name="category" class="form-control" value="{{.Topic.Category}}">
              </div>

              <button type="submit" class="btn btn-default">提交修改</button>
          </form>
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
