package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "mygo/routers"

)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	//orm.RegisterModel(new(Model))//注册 model
	orm.RegisterDataBase("", "mysql", "root:123@tcp(localhost:3306)/t9?charset=utf8", 30, 30) //注册默认数据库
}

func main() {
	beego.Run()
}

