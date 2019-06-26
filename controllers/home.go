package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//fmt.Println(checkAccount(c.Ctx))
	c.TplName = "home.html"
	topics, err := models.GetAllTopics(true)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["Topics"] = topics

}
