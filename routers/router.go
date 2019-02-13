package routers

import (
	"github.com/astaxie/beego"
	"go_code/blog_go/controllers"
)

func init() {
	beego.Router("/articles/", &controllers.ArticleController{})
	beego.Router("/articles/:id([0-9]+)/", &controllers.ArticleController{}, "get:GetDetail;delete:Delete")
}
