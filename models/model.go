package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stevechan1993/Beego-CMS/util"
)

func init() {
	driverName := beego.AppConfig.String("driverName")

	// 注册数据库驱动
	regError := orm.RegisterDriver(driverName, orm.DRMySQL)
	if regError != nil {
		util.LogError("注册驱动出错")
		return
	}

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
	orm.RegisterModel(
		new(Permission),
		new(City),
		new(FoodCategory),
		new(OrderStatus),
		new(Admin),
		new(User),
		new(Food),
		new(Shop),
		new(UserOrder),
		new(SupportService),
		new(Address))

	// 创建表
	orm.RunSyncdb("default", false, true)
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
	Id 				int 			`json:"id"`   			// 用户编号id
	UserName 		string  		`json:"username"`  		// 用户名称
	RegisterTime 	string 			`json:"register_time"`  // 用户注册时间
	Mobile 			string 			`json:"mobile"`  		// 用户手机号
	IsActive 		int 			`json:"is_active"`   	// 用户是否激活
	Balance 		int 			`json:"balance"`  		// 用户的账户余额
	Avatar 			string 			`json:"avatar"`  		// 用户头像
	City 			*City 			`orm:"rel(fk)"`  		// orm映射  用户所在城市 一对一关系 一个用户能有一个城市地区
	UserOrder 		[]*UserOrder 	`orm:"reverse(many)"`  	// orm映射 用户订单 一个用户可以有多张订单，设置一对多关系
	Pwd 			string 			`json:"password"`  		// 用户的账密码
	DelFlag 		int 			`json:"del_flag"`  		// 是否被删除的标志字段 软删除
}

// 食品种类表
type FoodCategory struct {
	Id 					int 	`json:"id"`    							// 食品id
	CategoryName 		string 	`json:"name" orm:"size(32)"`    		// 食品种类名称
	CategoryDesc 		string 	`json:"description" orm:"size(200)"` 	// 食品种类描述
	Level 				int 	`json:"level"` 							// 食品种类层级
	ParentCategoryId 	int 	`json:"parent_category_id"`   			// 父一级的类型id
	Restaurant 			*Shop 	`json:"restaurant_id" orm:"rel(fk)"` 	// 该食品种类所属的商铺id
	Food 				[]*Food `orm:"reverse(many)"`  					// 食品
}

// 食品表
type Food struct {
	Id 			int 			`json:"item_id"`   		// 食品id
	Name 		string 			`json:"name"`  			// 食品名称
	Description string 			`json:"description"`   	// 食品描述
	Rating 		int 			`json:"rating"`  		// 食品评分
	MonthSales 	int 			`json:"month_sales"`  	// 月销量
	ImagePath 	string 			`json:"image_path"` 	// 食品图片地址
	Activity 	string 			`json:"activity"`		// 食品活动
	Attributes 	string 			`json:"attributes"` 	// 食品特点
	Specs 		string 			`json:"specs"`  		// 食品规格
	Category 	*FoodCategory 	`orm:"rel(fk)"`  		// 食品种类
	Restaurant 	*Shop 			`orm:"rel(fk)"`  		// 食品店铺信息
	DelFlag 	int 			`json:"del_flag"`  		// 是否已经被删除 0表示未删除 1表示被删除
 }

// 商家店铺表
type Shop struct {
	Id 							int 				`json:"item_id"`  							// 店铺id
	Name 						string 				`json:"name"` 								// 店铺名称
	Address 					string 				`json:"address"` 							// 店铺地址
	Latitude 					float32 			`json:"latitude"` 							// 经度
	Longitude 					float32 			`json:"longitude"`  						// 纬度
	Description 				string 				`json:"description"` 						// 店铺简介
	Phone 						int64 				`json:"phone"` 								// 店铺电话
	PromotionInfo 				string 				`json:"promotion_info"` 					// 店铺标语
	FloatDeliveryFee 			int 				`json:"float_delivery_fee"`  				// 配送费
	FloatMinimumOrderAmount 	int 				`json:"float_minimum_order_amount"`  		// 起送价
	IsPremium 					bool 				`json:"is_premium"` 						// 品牌保障
	DeliveryMode 				bool 				`json:"delivery_mode"` 						// 配送方式
	New 						bool 				`json:"new"`  								// 新开店铺
	Bao 						bool 				`json:"bao"`  								// 外卖宝
	Zhun 						bool 				`json:"zhun"` 								// 准时达
	Piao 						bool 				`json:"piao"`  								// 开发票
	StartTime 					string 				`json:"startTime"`  						// 营业开始时间
	EndTime 					string 				`json:"endTime"`  							// 营业结束时间
	ImagePath 					string 				`json:"image_path"`  						// 店铺头像
	BusinessLicenseImage 		string 				`json:"business_license_image"` 			// 营业执照
	CateringServiceLicenseImage string 				`json:"catering_service_license_image"`  	// 餐饮服务许可证
	Category 					string 				`json:"category"` 							// 店铺类型
	RecentOrderNum 				int 				`json:"recent_order_num"`  					// 最近一个月的销量
	Rating 						int 				`json:"rating"`  							// 综合评分
	Activities 					[]*SupportService 	`json:"activities" orm:"reverse(many)"` 	// 一个商家对应多家服务
	UserOrder 					[]*UserOrder 		`json:"user_order" orm:"reverse(many)"` 	// 设置一对多关系： 一个店铺，可能会有多个订单
	Foods 						[]*Food 			`orm:"reversed(many)"` 						// 设置一对多关系的反向关系
	Dele 						int 				`json:"dele"` 								// 删除标志 1表示已经删除 0表示未删除
}

// 订单状态表
type OrderStatus struct {
	Id 			int
	StatusId 	int						  				// 订单状态编号
	StatusDesc 	string 			`orm:"size(100)"`   	// 订单状态描述
	UserOrder 	[]*UserOrder 	`orm:"reverse(many)"`  	// 一个订单状态可以对应多个订单
}

// 商家所支持的服务表
type SupportService struct {
	Id 			int 						// 编号
	Name 		string 						// 名称
	IconName 	string 						// 前端设置的图标内容
	IconColor 	string 						// 前端设置的图标颜色
	Description string 						// 服务描述
	Shop 		[]*Shop `orm:"rel(m2m)"`  	// orm映射 一个活动服务可以被多个商家参加
}

// 用户订单表
type UserOrder struct {
	Id 			int 			`json:"id"`  		// 用户订单编号id
	SumMoney 	int 			`orm:"default(0)"` 	// 用户订单总价格
	Time 		string 			`json:"time"`  		// 订单创建时间
	OrderTime 	uint64 			`json:"order_time"` // 订单创建时间
	OrderStatus *OrderStatus 	`orm:"rel(fk)"`  	// 设置一对一的关系，一张订单只能有一个状态
	User 		*User 			`orm:"rel(fk)"`  	// 一张订单只能对应一个用户
	Shop 		*Shop 			`orm:"rel(fk)"`  	// 设置一对一的关系： 一张订单只能有一个商家
	Address 	*Address 		`orm:"rel(fk)"`  	// 用户订单地址
	DelFlag 	int 			`json:"del_flag"` 	// 删除标志，软删除
}

// 订单地址表
type Address struct {
	Id 				int 			`json:"id"`  			// 订单地址id
	Address 		string 			`json:"address"`  		// 地址
	Phone 			string 			`json:"phone"`  		// 联系人手机号
	AddressDetail 	string 			`json:"address_detail"` // 地址详情
	IsValid 		int 			`json:"is_valid"`
	UserOrder 		[]*UserOrder 	`orm:"reverse(many)"`
}