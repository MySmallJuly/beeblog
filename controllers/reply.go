package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	tid := c.Input().Get("tid")
	err := models.AddReply(tid,c.Input().Get("nickname"),c.Input().Get("content"))
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/topic/view/"+tid,302)
}

func (c *ReplyController) Delete() {
	tid := c.Input().Get("tid")
	if !checkAccount(c.Ctx) {
		c.Redirect("/login",302)

		return
	}
	err := models.DeleteReply(c.Input().Get("rid"))
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/topic/view/"+tid,302)
}