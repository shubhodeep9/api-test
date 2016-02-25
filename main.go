package main

import (
	"api-test/Godeps/_workspace/src/github.com/astaxie/beego"
	"api-test/Godeps/_workspace/src/github.com/astaxie/beego/plugins/auth"
	_ "api-test/docs"
	_ "api-test/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("shubho", "deep"))
	beego.Run()
}
