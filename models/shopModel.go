package models

/**
 * 数据映射为json
 */
func (this *Shop) ShopToRespDesc() interface{} {
	respDesc := map[string]interface{} {
		"id":	this.Id,
		"name": this.Name,
		"address": this.Address,
		"phone": this.Phone,
		"status": this.Status,
		"recent_order_num": this.RecentOrderNum,
		"rating_count": this.RatingCount,
		"rating": this.Rating,
	}
	return respDesc
}