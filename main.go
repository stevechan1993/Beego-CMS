package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/stevechan/Beego-CMS/routers"
)

func main() {
	beego.Run()
}

