{{template "header"}}
        <title>登录 - 我的beego博客</title>
    </head>

    <body>
        <div class="container" style="width: 500px">
            <form method="POST" action="/login">
              <div class="form-group">
                <label>Account</label>
                <input id="uname" class="form-control" placeholder="Enter account" name="uname">
              </div>
              <div class="form-group">
                <label>Password</label>
                <input id="pwd" type="password" class="form-control" placeholder="Password" name="pwd">
              </div>
              <div class="checkbox">
                <label>
                <input type="checkbox" name="autologin"> 自动登录
                </label>
              </div>
              <button type="submit" class="btn btn-default" onclick="return checkInput();">登录</button>
              <button type="submit" class="btn btn-default" onclick="return backToHome();">返回首页</button>
            </form>

            <script type="text/javascript">
                function checkInput() {
                  var uname = document.getElementById("uname");
                  if (uname.value.length == 0) {
                    alert("请输入账号");
                    return false;
                  }

                  var pwd = document.getElementById("pwd");
                  if (pwd.value.length == 0){
                    alert("请输入密码");
                    return false;
                  }

                  return true;

                }

                function backToHome() {
                  window.location.href = "/";
                  return false;
                }
          </script>
        </div>
    </body>
</html>
