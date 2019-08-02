package main

import (
	//_ "github.com/astaxie/beego/cache/redis"
	//"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego"
	_ "github.com/yijie8/tmcapi/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		//bm, err := cache.NewCache("redis", `{"key":"tmc","conn":"117.78.46.158:6000","dbNum":"5","password":""}`)
		//cache.Test()
	}
	beego.Run()
}
