package models

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (this *User) UserToRespDesc() interface{} {
	respInfo := map[string] interface{} {
		"id":			this.Id,
		"user_id":		this.Id,
		"username":		this.UserName,
		"city":			this.City.CityName,
		"register_time":this.RegisterTime,
		"mobile":		this.Mobile,
		"is_active": 	this.IsActive,
		"balance":		this.Balance,
		"avatar":		this.Avatar,
	}
	return respInfo
}