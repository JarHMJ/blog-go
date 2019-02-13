package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "go_code/blog_go/models"
	_ "go_code/blog_go/routers"
	"os"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)

}

func migrate() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "migrate" {
		migrate()
		return
	}
	beego.Run()
}
