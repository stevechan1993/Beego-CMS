package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/stevechan1993/Beego-CMS/models"
	"github.com/stevechan1993/Beego-CMS/util"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	beego.Controller
}

/**
 * 更新用户头像
 */
func (this *FileController) UpdateAdminAvatar() {

	util.LogInfo("更新用户头像")

	resp := make(map[string] interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 判断是否登录的权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.ERROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.ERROR_UNLOGIN)
		return
	}

	// 获取文件操作
	file, head, err := this.GetFile("file")
	defer file.Close()

	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	fileArr := strings.Split(head.Filename, ".")
	// 文件类型判断
	if (fileArr[1] != "png" && fileArr[1] != "jpg" && fileArr[1] != "jpeg") {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURETYPE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURETYPE)
		return
	}

	// 文件大小判断，控制文件在2M以内
	if head.Size > 1024*1024*2 {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURESIZE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURESIZE)
		return
	}

	uploadPath := "./img/"
	// 判断upload是否存在，不存在先创建
	if exist, _ := util.IsExists(uploadPath); !exist {
		if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
			beego.Info(err.Error())
			resp["status"] = util.RECODE_FAIL
			resp["type"] = util.RESPMSG_ERROR_PICTUREADD
			resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
			return
		}
	}

	// 目录创建成功，继续保存文件
	fileArray := strings.Split(head.Filename, ".")
	fileName := "avatar" + strconv.Itoa(int(time.Now().UnixNano())) + "." + fileArray[1]
	path := uploadPath + fileName

	if err = this.SaveToFile("file", path); err != nil {
		beego.Info(err.Error())
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	// 文件保存到目录成功，更新数据库
	om := orm.NewOrm()
	adminId, _ := strconv.Atoi(this.Ctx.Input.Param(":adminId"))
	admin := models.Admin{Id: adminId}
	if om.Read(&admin) == nil {
		admin.Avatar = fileName
		if _, err := om.Update(&admin, "avatar"); err == nil {
			resp["status"] = util.RECODE_OK
			resp["image_path"] = fileName
			return
		}
	}

	resp["status"] = util.RECODE_OK
	resp["type"] = util.RESPMSG_ERROR_PICTUREADD
	resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
}

/**
 * 上传图片
 */
func (this *FileController) UploadImg() {

	util.LogInfo("上传图片")

	resp := make(map[string] interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	// 使用getFile来获取上传的文件
	// file为文件类型结构体，head可以得到文件的大小和文件的名字
	file, head, err := this.GetFile("file")
	beego.Info("文件：", head.Filename, head.Size)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	fileArr := strings.Split(head.Filename, ".")
	// 文件类型判断
	if (fileArr[1] != "png" && fileArr[1] != "jpg" && fileArr[1] != "jpeg") {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURETYPE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURETYPE)
		return
	}

	// 文件大小判断，控制文件在2M以内
	if head.Size > 1024*1024*2 {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURESIZE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURESIZE)
		return
	}

	// 切记关闭文件
	path := "./img/" + head.Filename
	file.Close()
	err = this.SaveToFile("file", path)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	// 返回正确的图片链接
	resp["status"] = util.RECODE_OK
	resp["image_path"] = head.Filename
}

/**
 * 判断用户是否已经登录过：通过session进行判断
 */
func (this *FileController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}