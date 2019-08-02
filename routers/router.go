// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/yijie8/tmcapi/controllers"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		if !ctx.Input.IsAjax() && ctx.Input.Param(":test")=="" {
			jsonData := make(map[string]interface{}, 3)
			jsonData["Code"] = 403
			jsonData["Message"] = "必须为ajax请求"
			returnJSON, _ := json.Marshal(jsonData)
			_, _ = ctx.ResponseWriter.Write(returnJSON)
		}
	}
	// 添加过滤器 不是ajax的请求都过滤掉
	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v2/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/v3/*", beego.BeforeRouter, FilterUser)


	var FinishRouter = func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Add("TMCAPI-Version", "1.0.8")
		ctx.ResponseWriter.Header().Add("X-XSS-Protection", "1; mode=block")
	}

	beego.InsertFilter("*", beego.BeforeRouter, Allow(&Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, FinishRouter, false)
	//唯一请求 验证UToken在30分钟内是否唯一







	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	//beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
