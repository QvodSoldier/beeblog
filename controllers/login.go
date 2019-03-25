package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit")
	beego.Info(isExit)
	if isExit == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}

	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autologin") == "on"

	if uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd") {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	uname := ctx.GetCookie("uname")
	if uname == "" {
		return false
	}

	pwd := ctx.GetCookie("pwd")
	if pwd == "" {
		return false
	}

	return uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd")
}
