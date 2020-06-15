package routers

import (
	"github.com/astaxie/beego"
	"github.com/stevechan/beego-demo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
