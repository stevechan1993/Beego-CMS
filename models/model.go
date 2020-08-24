package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stevechan/Beego-CMS/util"
)

func init() {
	driverName := beego.AppConfig.String("driverName")

	// 注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	// 数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		util.LogError("连接数据库出错")
		return
	}
	util.LogInfo("连接数据库成功")

	// 注册实体模型
	orm.RegisterModel()

	// 创建表
}

// 管理员权限等级及级别名称
type Permission struct {
	Id		int		                               // 权限登记id
	Level	string	`json:"level" orm:"size(30)"`  // 权限级别
	Name	string	`json:"name" orm:"size(20)"`   // 权限名称
	Admin	[]*Admin `orm:rel(m2m)`                // orm映射 一个权限可以被多个管理员所拥有
}

// 地区城市表
type City struct {

}

// 管理员表
type Admin struct {

}

// 用户信息表
type User struct {

}

// 食品种类表
type FoodCategory struct {

}

// 食品表
type Food struct {

}

// 商家店铺表
type Shop struct {

}

// 订单状态表
type OrderStatus struct {

}

// 商家所支持的服务表
type SupportService struct {

}

// 用户订单表
type UserOrder struct {

}

// 订单地址表
type Address struct {

}