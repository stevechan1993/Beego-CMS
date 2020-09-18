package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/stevechan/Beego-CMS/models"
	"github.com/stevechan/Beego-CMS/util"
	//"github.com/stevechan/Beego-CMS/entity"
	//"github.com/astaxie/beego/orm"
	//"strings"
	//"math/rand"
	//"time"
)

type AdminController struct {
	beego.Controller
}

const (
	ADMINTABLENAME = "admin"
	ADMIN = "admin"
)

/**
管理员登录
 */
func (this *AdminController) AdminLogin() {

}

/**
获取管理员信息
 */
func (this *AdminController) GetAdminInfo() {

	util.LogInfo("获取管理员信息")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	userByte := this.GetSession(ADMIN)
	if userByte == nil {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.ERROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}
	var admin models.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)
	if err != nil {
		util.LogInfo("获取管理员信息失败")
		reJson["status"] = util.RECODE_FAIL
		reJson["type"] = util.RESPMSG_ERRORSESSION
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERRORSESSION)
		return
	}
	if (admin.Id > 0) {
		util.LogInfo("获取管理员信息成功")
		reJson["status"] = util.RECODE_OK
		reJson["data"] = admin.AdminToRespDesc()
		return
	}
}

/**
退出登录
 */
func (this *AdminController) SignOut() {

	util.LogInfo("管理员退出当前账号")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 删除session
	this.DelSession(ADMIN)

	resp["status"] = util.RECODE_OK
	resp["success"] = util.Recode2Text(util.RESPMSG_SIGNOUT)
}

/**
获取管理员总数
 */
func (this *AdminController) GetAdminCount() {

}

/**
返回管理员当日统计结果
 */
func (this *AdminController) GetAdminStatis() {

}

/**
获取管理员列表
 */
func (this *AdminController) GetAdminList() {
	util.LogInfo("管理员列表")
	reJSon := make(map[string]interface{})
	this.Data["json"] = reJSon
	defer this.ServeJSON()
	if !this.IsLogin() {
		reJSon["status"] = util.RECODE_UNLOGIN
		reJSon["type"] = util.ERROR_UNLOGIN
		reJSon["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
	}
}

/**
判断用户是否已经登录过：通过session进行判断
 */
func (this *AdminController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}