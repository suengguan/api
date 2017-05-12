package main

import (
	_ "api/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Debug("this is api service v1")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
