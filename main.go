package main

import (
	"github.com/astaxie/beego"
	_ "github.com/stevechan1993/Beego-CMS/routers"
)

func main() {

	beego.AddAPPStartHook(func() error {
		beego.Info("自定义配置")
		return nil
	})

	// 静态资源路径映射
	beego.SetStaticPath("manage/static", "static")

	// 图片资源路径映射
	beego.SetStaticPath("img", "img")

	// 监听
	beego.Run()
}