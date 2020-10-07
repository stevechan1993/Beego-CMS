package controllers

import (
	"Beego-CMS/entity"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/stevechan1993/Beego-CMS/models"
	"github.com/stevechan1993/Beego-CMS/util"
	"math/rand"
	"strings"
	"time"
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

	// 先获取redis缓存实例
	redisConn, err := util.GetRedis()
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
		return
	}

	paths := strings.Split(this.Ctx.Input.URL(), "/")

	// 从redis中根据key值获取对应的数据缓存
	statis := redisConn.Get("userStatis" + paths[3])

	if statis != nil {
		var statisCount entity.StatisEntity
		json.Unmarshal(statis.([]byte), &statisCount)
		resp["status"] = util.RECODE_FAIL
		resp["count"] = statisCount.StaticCount
		return
	}

	om := orm.NewOrm()

	userCount, err := om.QueryTable(USERTABLENAME).Count()
	if err != nil {
		beego.Info(userCount)
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
		return
	}

	todayStr := time.Now().Format("2006-01-02")
	statisCount := &entity.StatisEntity{
		StaticDate: "userStatis" + paths[3],
		StaticCount: rand.Intn(150),
	}
	bytes, _ := json.Marshal(statisCount)

	if todayStr == paths[3] {
		redisConn.Put("userStatis" + paths[3], bytes, 60*time.Second)
	} else {
		redisConn.Put("userStatis" + paths[3], bytes, 60*60*24*time.Second)
	}

	resp["status"] = util.RECODE_OK
	resp["count"] = statisCount.StaticCount
}

/**
 * 获取总用户数
 */
func (this *UserController) GerUserCount() {
	util.LogInfo("获取用户总数")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 判断是否已经登录，未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.ERROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	userCount, err := om.QueryTable(USERTABLENAME).Count()
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["count"] = 0
	} else {
		resp["status"]= util.RECODE_OK
		resp["count"] = userCount
	}
}

/**
 * 获取用户信息列表
 */
func (this *UserController) UserList() {
	util.LogInfo("获取用户列表")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 判断是否已经登录，未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"]= util.ERROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	// 查询数据
	var userList []*models.User
	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	om := orm.NewOrm()
	om.QueryTable(USERTABLENAME).Filter("del_flag", 0).Limit(limit, offset).All(&userList)

	// 使用loadRelated方法进行关联查询，并进行json格式组装
	var respList []interface{}
	for _, user := range userList {
		om.LoadRelated(user, "City")
		respList = append(respList, user.UserToRespDesc())
	}

	// 返回查询数据
	if len(userList) > 0 {
		this.Data["json"] = &respList
	} else {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_USERLIST
		resp["message"]= util.Recode2Text(util.RESPMSG_ERROR_USERLIST)
	}
}

/**
 * 通过用户名查询用户信息
 */
func (this *UserController) GerUserInfoByUserName() {
	util.LogInfo("通过用户名查询用户信息")

	resp := make(map[string]interface{})
	defer this.ServeJSON()

	// 判断是否已经登录，未登录，返回没有权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"]= util.ERROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	om := orm.NewOrm()
	var user models.User
	om.QueryTable(USERTABLENAME).Filter("user_name", this.Ctx.Input.Param(":username")).One(&user)

	if user.Id <= 0 {
		resp["status"]= util.RECODE_FAIL
		resp["type"]= util.RESPMSG_ERROR_USERINFO
		resp["message"] = util.Recode2Text(util.RESPMSG_ERROR_USERINFO)
		this.Data["json"]= resp
		return
	}
	this.Data["json"] = user.UserToRespDesc()
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