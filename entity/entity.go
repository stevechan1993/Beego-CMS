package entity

/**
用户登录接口实体
 */
type AdminLoginEntity struct {
	User_name string `json:"user_name"`
	Password string `json:"password"`
}

/**
某一日期统计的结果
 */
type StatisEntity struct {
	StaticDate string `json:"statis_date"`
	StaticCount int `json:"statis_count"`
}

/**
添加食品信息
 */
type AddFoodEntity struct {

}