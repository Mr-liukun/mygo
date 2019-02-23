package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mygo/models"
)

type MainController struct {
	beego.Controller

}

func (c *MainController) One(){
	fmt.Println("进来看了")
	var so models.MysqlModel
	so = models.SelectDB()

	// 向页面传递参数
	c.Data["id"] = so.Id
	c.Data["name"] = so.Name
	// 跳转到view/one.html文件
	c.TplName = "one.html"
}
