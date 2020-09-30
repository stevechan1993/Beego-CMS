package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/stevechan1993/Beego-CMS/models"
	"github.com/stevechan1993/Beego-CMS/util"
	"strings"
)

type FileController struct {
	beego.Controller
}

/**
 * 更新用户头像
 */
func (this *FileController) UpdateAdminAvatar() {

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