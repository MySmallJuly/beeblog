package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	op := c.Input().Get("op")
	switch op {
	case "add":
		name :=c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect("/category",302)
		return
	case "del":
		id :=c.Input().Get("id")
		fmt.Println(id+"444444")
		fmt.Println("__________")
		if len(id) == 0 {
			break
		}
		err := models.DelAllCategories(id)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect("/category",302)
		return
	}
	c.Data["IsCategory"] = true
	c.Data["Categories"],_ = models.GetAllCategories()
	c.TplName = "category.html"
}
