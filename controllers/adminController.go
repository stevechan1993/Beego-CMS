package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/stevechan1993/Beego-CMS/entity"
	"github.com/stevechan1993/Beego-CMS/models"
	"github.com/stevechan1993/Beego-CMS/util"
	"math/rand"
	"strings"
	"time"
)

type AdminController struct {
	beego.Controller
}

const (
	ADMINTABLENAME = "admin"
	ADMIN = "admin"
)

/**
 * 管理员登录
 */
func (this *AdminController) AdminLogin() {
	util.LogInfo("管理员登录")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	// 获取请求数据
	var loginEntity entity.AdminLoginEntity
	util.JsonToEntity(this.Ctx.Input.RequestBody, &loginEntity)

	// 查询结果变量
	var admin models.Admin

	// 实例化orm对象
	om := orm.NewOrm()

	// select * from admin where user_name = ? and pwd = ? values root, 1234
	om.QueryTable(ADMINTABLENAME).Filter("user_name", loginEntity.User_name).Filter("pwd", loginEntity.Password).One(&admin)

	// 管理员成功登录
	if (admin.Id > 0) {

		userByte, _ := json.Marshal(admin)

		// 设置session
		this.SetSession(ADMIN, userByte)

		reJson["status"] = util.RECODE_OK
		reJson["success"] = util.Recode2Text(util.RESPMSG_SUCCESSLOGIN)
		return
	}

	reJson["status"] = util.RECODE_FAIL
	reJson["message"] = util.Recode2Text(util.RESPMSG_FAILURELOGIN)
}

/**
 * 获取管理员信息
 */
func (this *AdminController) GetAdminInfo() {

	util.LogInfo("获取管理员信息")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	// 从session中获取信息
	userByte := this.GetSession(ADMIN)

	// 判断session是否为空
	if userByte == nil {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.ERROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	var admin models.Admin

	err := json.Unmarshal(userByte.([]byte), &admin)

	// 失败
	if err != nil {
		util.LogInfo("获取管理员信息失败")
		reJson["status"] = util.RECODE_FAIL
		reJson["type"] = util.RESPMSG_ERRORSESSION
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERRORSESSION)
		return
	}

	// 成功
	if (admin.Id > 0) {
		util.LogInfo("获取管理员信息成功")
		reJson["status"] = util.RECODE_OK
		reJson["data"] = admin.AdminToRespDesc()
		return
	}
}

/**
 * 退出登录
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
 * 获取管理员总数
 */
func (this *AdminController) GetAdminCount() {
	util.LogInfo("获取管理员总数")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	// 判断是否有权限
	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.ERROR_UNLOGIN
		reJson["messae"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	adminCount, err := om.QueryTable(ADMINTABLENAME).Filter("status", 0).Count()
	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERRORADMINCOUNT)
		reJson["count"] = 0
	} else {
		reJson["status"] = util.RECODE_OK
		reJson["count"] = adminCount
	}
}

// TODO
/**
 * 返回管理员当日统计结果
 */
func (this *AdminController) GetAdminStatis() {
	util.LogInfo("获取管理员某个日期统计结果")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	// 判断是否有权限的优化
	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.ERROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	// 获取redis缓存实例
	redisConn, err := util.GetRedis()
	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["count"] = 0
		return
	}

	// 获取参数
	paths := strings.Split(this.Ctx.Input.URL(), "/")

	// 从redis中根据key值获取对应的数据缓存
	statis := redisConn.Get("adminStatis" + paths[3])

	if statis != nil {
		var statisCount entity.StatisEntity
		json.Unmarshal(statis.([]byte), &statisCount)
		reJson["status"] = util.RECODE_OK
		reJson["count"] = statisCount.StaticCount
		return
	}

	om := orm.NewOrm()

	// 仅做测试效果使用
	adminCount, err := om.QueryTable(ADMINTABLENAME).Count()
	if err != nil {
		beego.Info(adminCount)
		reJson["status"] = util.RECODE_FAIL
		reJson["count"] = 0
		return
	}

	todayStr := time.Now().Format("2006-01-02")
	statisCount := &entity.StatisEntity{
		StaticDate: "adminStatis" + paths[3],
		StaticCount: rand.Intn(50),
	}
	bytes, _ := json.Marshal(statisCount)

	// 分类存储到redis中
	if todayStr == paths[3] {
		redisConn.Put("adminStatis" + paths[3], bytes, 60 * time.Second)
	} else {
		redisConn.Put("adminStatis" + paths[3], bytes, 60 * 60 * 24 * time.Second)
	}

	reJson["status"] = util.RECODE_OK
	reJson["count"] = statisCount.StaticDate
}

/**
 * 获取管理员列表
 */
func (this *AdminController) GetAdminList() {

	util.LogInfo("管理员列表")

	reJson := make(map[string]interface{})
	this.Data["json"] = reJson
	defer this.ServeJSON()

	if !this.IsLogin() {
		reJson["status"] = util.RECODE_UNLOGIN
		reJson["type"] = util.ERROR_UNLOGIN
		reJson["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	var adminList []*models.Admin
	om := orm.NewOrm()
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")
	_, err := om.QueryTable(ADMINTABLENAME).Filter("status", 0).Limit(limit, offset).All(&adminList)

	if err != nil {
		reJson["status"] = util.RECODE_FAIL
		reJson["type"] = util.RESPMSG_ERROR_FOODLIST
		reJson["message"] = util.Recode2Text(util.RESPMSG_ERROR_FOODLIST)
		return
	}

	var respList []interface{}
	for _, admin := range adminList {
		om.LoadRelated(admin, "City")
		respList = append(respList, admin.AdminToRespDesc())
	}

	reJson["status"] = util.RECODE_OK
	reJson["data"] = respList
}

/**
 * 判断用户是否已经登录过：通过session进行判断
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