package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/stevechan/beego-demo/routers"
)

func main() {
	beego.Run()
}

