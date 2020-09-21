package entity

/**
用户登录接口实体
 */
type AdminLoginEntity struct {
	User_name string `json:"user_name"`
	Password  string `json:"password"`
}

/**
某一日期统计的结果
 */
type StatisEntity struct {
	StaticDate 	string 	`json:"statis_date"`
	StaticCount int 	`json:"statis_count"`
}

/**
添加食品信息
 */
type AddFoodEntity struct {
	Name 			string 		`json:"name"`   		// 食品名称
	Description 	string 		`json:"description"`  	// 食品描述
	ImagePath 		string 		`json:"image_path"`  	// 食品图片地址
	Activity 		string 		`json:"activity"`  		// 食品活动
	Attributes 		[]string 	`json:"attributes"` 	// 食品特点
	Specs 			[]Specs 	`json:"specs"` 			// 食品规格
	CategoryId 		int 		`json:"category_id"` 	// 食品种类  种类id
	RestaurantId 	string 		`json:"restaurant_id"`  // 店铺id
}

/**
食品规格
 */
type Specs struct {
	Specs 		string 	`json:"specs"`
	PackingFee 	int 	`json:"packing_fee"`
	Price 		int 	`json:"price"`
}

/**
搜索地址信息的实体
 */
type PoiSearch struct {
	Name 		string 	`json:"name"`
	Address 	string 	`json:"address"`
	Latitude 	float32 `json:"latitude"`
	Longitude 	float32 `json:"longitude"`
	Geohash 	string 	`json:"geohash"`
}