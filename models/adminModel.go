package models

/**
从Admin数据库实体转换为前端请求resp的json格式
 */
func (admin *Admin) AdminToRespDesc() interface{} {
	respDesc := map[string]interface{} {
		"user_name": 	admin.UserName,
		"id": 			admin.Id,
		"create_time": 	admin.CreateTime,
		"status": 		admin.Status,
		"avatar": 		admin.Avatar,
		"city": 		admin.City.CityName,
		"admin": 		"管理员",
	}
	return respDesc
}