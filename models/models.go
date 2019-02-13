package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var db orm.Ormer

type Article struct {
	Id      int
	Name    string
	Content string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Tags    []*Tag    `orm:"rel(m2m)"`
}

type Tag struct {
	Id       int
	Name     string
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")
	orm.RegisterModel(new(Article), new(Tag))
	db = orm.NewOrm()
}

func GetDB() orm.Ormer {
	return db
}

func GetQuerySeter(table string) orm.QuerySeter {
	return db.QueryTable(table)
}
