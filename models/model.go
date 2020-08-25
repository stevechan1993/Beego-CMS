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
	Id		int		                               		// 权限登记id
	Level	string		`json:"level" orm:"size(30)"`  	// 权限级别
	Name	string		`json:"name" orm:"size(20)"`   	// 权限名称
	Admin	[]*Admin 	`orm:"rel(m2m)"`                // orm映射 一个权限可以被多个管理员所拥有
}

// 地区城市表
type City struct {
	Id 			int 		`json:"id"`  					//城市id
	CityName 	string 		`json:"name" orm:"size(20)"`   	//城市名称
	PinYin 		string 		`json:"pin_yin"`  				//城市名称拼音
	Longitude 	float32 	`json:"longitude"` 	 			// 城市经度
	Latitude 	float32 	`json:"latitude"`  				// 城市纬度
	AreaCode 	string 		`json:"area_code"`   			// 城市地区编码
	Abbr		string 		`json:"abbr"`    				// 城市拼音缩写
	User 		[]*User 	`orm:"reverse(many)"`   		// orm映射 一个城市可以有多个用户
	Admin 		[]*Admin 	`orm:"reverse(many)"`   		// orm映射 一个城市可以有多个管理员
}

// 管理员表
type Admin struct {
	Id 			int 			`json:"id"`   						// 管理员编号id
	UserName 	string 			`json:"user_name" orm:"size(12)"`   // 管理员用户名
	CreateTime 	string 			`json:"create_time"`   				// 记录添加时间
	Status 		int 			`json:"status"`  					// 管理员状态
	Avatar 		string			`json:"avatar" orm:"size(50)"`  	// 管理员头像
	Pwd 		string 			`json:"pwd"`   						// 管理员密码
	Permission 	[]*Permission 	`orm:"reverse(many)"`   			// 一个管理员可以有多种权限
	City 		*City 			`orm:"rel(fk)"`  					// orm映射 管理员所在城市
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