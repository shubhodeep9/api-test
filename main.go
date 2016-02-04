package main

import (
	_ "api-test/docs"
	_ "api-test/routers"
	"github.com/astaxie/beego/plugins/auth"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("shubho","deep"))
	beego.Run()
}
