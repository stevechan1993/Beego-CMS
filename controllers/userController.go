package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/stevechan1993/Beego-CMS/models"
	"github.com/stevechan1993/Beego-CMS/util"
)

type UserController struct {
	beego.Controller
}

const (
	USERTABLENAME = "user"
)

/**
 * 获取某日新增用户
 */
func (this *UserController) UserStatisDaily() {
	util.LogInfo("获取用户某日增长量统计")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 判断是否已经登录，是否有查询权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.ERROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}


}

/**
 * 获取总用户数
 */
func (this *UserController) GerUserCount() {

}

/**
 * 获取用户信息列表
 */
func (this *UserController) UserList() {

}

/**
 * 通过用户名查询用户信息
 */
func (this *UserController) GerUserInfoByUserName() {

}

/**
 * 判断用户是否已经登录
 */
func (this *UserController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		_ = json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}