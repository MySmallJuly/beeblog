package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get()  {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["Topics"] = topics
}


func (c *TopicController) Post()  {
	fmt.Println(checkAccount(c.Ctx))
	if !checkAccount(c.Ctx) {
		c.Redirect("/login",302)

		return
	}
	tid := c.Input().Get("id")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	var err error
	if len(tid) == 0  {
		err = models.AddTopic(title,content,category)
	} else {
		err = models.ModifyTopic(tid,title,content,category)
	}

	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/topic",302)
}

func (c *TopicController) Add()  {
	c.TplName = "topic_add.html"
}



func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	tid := c.Ctx.Input.Param("0")
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
	fmt.Println(c.Ctx.Input.Param("0"))
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"
	tid :=c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err  != nil {
		c.Redirect("/",302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["tid"] = tid
}


func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login",302)
		return
	}
	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/",302)
}