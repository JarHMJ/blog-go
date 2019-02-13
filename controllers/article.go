package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_code/blog_go/models"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	var articles []models.Article
	models.GetQuerySeter("article").All(&articles)
	for _, v := range articles {
		fmt.Print(v)
	}
	c.Data["json"] = articles
	c.ServeJSON()
}

func (c *ArticleController) GetDetail() {
	id := c.Ctx.Input.Param(":id")
	var article models.Article
	err := models.GetQuerySeter("article").Filter("id", id).One(&article)
	if err == orm.ErrNoRows {
	} else {
		c.Data["json"] = article
	}
	c.ServeJSON()
}

func (c *ArticleController) Delete() {
	id := c.Ctx.Input.Param(":id")
	var article models.Article
	err := models.GetQuerySeter("article").Filter("id", id).One(&article)
	if err == orm.ErrNoRows {
	} else {
		c.Data["json"] = article
	}
	c.ServeJSON()
}
