package routers

import (
	"github.com/astaxie/beego"
	"github.com/stevechan/Beego-CMS/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    /**管理员相关操作**/
    beego.Router("/admin/info", &controllers.AdminController{}, "GET:GetAdminInfo")
    beego.Router("/admin/signout", &controllers.AdminController{}, "GET:SignOut")
    beego.Router("/admin/all", &controllers.AdminController{}, "GET:GetAdminList")
}
