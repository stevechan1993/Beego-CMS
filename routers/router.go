package routers

import (
	"github.com/astaxie/beego"
	"github.com/stevechan1993/Beego-CMS/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    /**************************************管理员相关操作******************************************/
    beego.Router("/admin/info", &controllers.AdminController{}, "GET:GetAdminInfo")   // 登录
    beego.Router("/admin/signout", &controllers.AdminController{}, "GET:SignOut")  // 获取管理员信息
    beego.Router("/admin/all", &controllers.AdminController{}, "GET:GetAdminList") // 登出
    beego.Router("/admin/count", &controllers.AdminController{}, "GET:GetAdminCount")  // 获取管理员总数
    beego.Router("/statis/admin/*/count", &controllers.AdminController{}, "GET:GetAdminStatis")  // 获取某一日的管理员增长统计数据
    beego.Router("/admin/all", &controllers.AdminController{}, "GET:GetAdminList")  // 查询所有的用户列表

    /*************************************用户模块相关操作****************************************/

    /*************************************商家店铺相关操作****************************************/

    /*************************************食品种类模块操作****************************************/

    /*************************************食品模块相关操作****************************************/

    /*************************************订单模块相关操作****************************************/

    /*************************************文件相关操作*******************************************/
    beego.Router("/admin/update/avatar/:adminId", &controllers.FileController{}, "POST:UpdateAdminAvatar")
    beego.Router("/v1/addimg/:username", &controllers.FileController{}, "POST:UploadImg")
}
