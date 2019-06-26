package models

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME = "mysql"
	_DRIVER  = "mysql"
)

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopiclastUserId int64
}

//文章
type Topic struct {
	Id          int64
	UID         int64
	Title       string
	Content     string `orm:"size(5000)"`
	Category    string
	Attachment  string    //附件
	Created     time.Time `orm:"index"`
	Updated     time.Time `orm:"index"`
	Views       int64     `orm:"index"`
	Author      string
	ReplyTime   time.Time `orm:"index"`
	ReplyCount  int64
	ReplyLastId int64
}

//评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic),new(Comment))
	orm.RegisterDriver(_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _DRIVER, "root:qianjiayu@/beeblog?charset=utf8")

}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		fmt.Println(err)
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelAllCategories(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid, Created: time.Now(), TopicTime: time.Now()}
	_, err = o.Delete(cate)
	fmt.Println(err)
	return err
}

func AddTopic(title, content, category string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	fmt.Println(err)
	return err

}

func GetAllTopics(isDics bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDics {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)

	}
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")

	err = qs.Filter("id", tidNum).One(topic)

	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, content, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: tidNum,
	}
	if o.Read(topic) == nil {
		topic.Category = category
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()

	}
	_, err = o.Update(topic)
	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: tidNum,
	}
	_, err = o.Delete(topic)
	return err
}

func AddReply(tid,nickname,content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:tidNum,
		Name:nickname,
		Content:content,
		Created:time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return err
}

func GetAllReplies(tid string) ([]*Comment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil,err
	}
	replies := make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_,err = qs.Filter("tid", tidNum).All(&replies)
	return replies,err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	reply := &Comment{
		Id:ridNum,
	}
	_, err = o.Delete(reply)
	return err
}