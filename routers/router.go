package routers

import (
	"github.com/astaxie/beego"
	"github.com/stevechan/Beego-CMS/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
