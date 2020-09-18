package models

/**
从Admin数据库实体转换为前端请求resp的json格式
 */
func (this *Admin) AdminToRespDesc() interface{} {

	respDesc := map[string]interface{} {
		"user_name": this.UserName,
		"id": this.Id,
		"create_time": this.CreateTime,
		"status": this.Status,
		"avatar": this.Avatar,
		"city": this.City.CityName,
		"admin": "管理员",
	}
	return respDesc
}