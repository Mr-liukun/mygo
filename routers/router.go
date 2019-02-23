package routers

import (
	"github.com/astaxie/beego"
	"mygo/controllers"
)

func init() {
    beego.Router("/mygo/one", &controllers.MainController{},"get:One")
}
