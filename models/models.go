package models

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beelog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm: "indexs"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Labels          string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddTopic(title, category, content, label string) error {
	o := orm.NewOrm()
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	topic := &Topic{
		Title:    title,
		Category: category,
		Labels:   label,
		Content:  content,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	// 改变分类数
	categories := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(categories)
	if err != nil {
		return err
	}
	categories.TopicCount++
	_, err = o.Update(categories)
	return err
	// 结束改变分类数
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}
	topic.ReplyCount++
	topic.ReplyTime = time.Now()
	_, err = o.Update(topic)
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}
	beego.Info(cate)

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("Comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

func GetAllTopics(cate, label string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("Topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

// func ChangeCategoriesNums(category, op string) (*Category, error) {
// 	o := orm.NewOrm()
//
// 	categories := new(Category)
//
// 	qs := o.QueryTable("category")
// 	err := qs.Filter("title", category).One(categories)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// if categories != nil {
// 	// 	categories.TopicCount++
// 	// 	_, err = o.Update(categories)
// 	// }
// 	//
// 	return categories, nil
// }

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topics := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topics)
	if err != nil {
		return nil, err
	}

	topics.Views++
	_, err = o.Update(topics)

	topics.Labels = strings.Replace(strings.Replace(
		topics.Labels, "#", " ", -1), "$", "", -1)
	return topics, err
}

func ModifyTopic(tid, title, category, content, label string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	// 改变分类数
	oldTopic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(oldTopic)
	if err != nil {
		return err
	}

	// 暂停
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Labels = label
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	// 继续

	if oldTopic.Category == category {
		beego.Info("didn't modify category")
		return nil
	}

	oldCategories := new(Category)
	newCategories := new(Category)
	qs2 := o.QueryTable("category")
	err = qs2.Filter("title", oldTopic.Category).One(oldCategories)
	if err != nil {
		return err
	}
	err = qs2.Filter("title", category).One(newCategories)
	if err != nil {
		return err
	}

	oldCategories.TopicCount--
	o.Update(oldCategories)
	newCategories.TopicCount++
	o.Update(newCategories)
	//结束该别分类数
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	// qs := o.QueryTable("category")
	// err = qs.Filter("id", cid).One(cate)
	// if err == nil {
	// 	return err
	// }

	_, err = o.Delete(cate)
	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	// 改变分类数
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}
	categories := new(Category)
	qs2 := o.QueryTable("category")
	err = qs2.Filter("title", topic.Category).One(categories)
	if err != nil {
		return err
	}
	categories.TopicCount--
	o.Update(categories)
	// 结束改变分类数
	_, err = o.Delete(topic)
	return err
}

func DeleteReply(tid, rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	reply := &Comment{Id: ridNum}
	_, err = o.Delete(reply)
	if err != nil {
		return err
	}

	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}

	replices := make([]*Comment, 0)
	qs2 := o.QueryTable("comment")
	_, err = qs2.Filter("tid", tidNum).OrderBy("-created").All(&replices)
	if err != nil {
		return err
	}

	topic.ReplyCount--
	if len(replices) != 0 {
		topic.ReplyTime = replices[0].Created
	}
	// topic.ReplyTime = nil
	o.Update(topic)

	return nil
}
