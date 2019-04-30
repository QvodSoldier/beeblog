package controllers

import (
	"fmt"

	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "index.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	fmt.Println(c.Input().Get("cate"))
	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
	fmt.Println(topics)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Categories"] = categories
	}
}
