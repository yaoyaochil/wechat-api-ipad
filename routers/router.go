// @APIVersion 7.0.12
// @Title Wechat
// @Description 仅供测试
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"wechatwebapi/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/712",
		beego.NSNamespace("/Api/Login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/Api/Message",
			beego.NSInclude(
				&controllers.MsgController{},
			),
		),
		beego.NSNamespace("/Api/Friend",
			beego.NSInclude(
				&controllers.FriendController{},
			),
		),
		// beego.NSNamespace("/Finder",
		// 	beego.NSInclude(
		// 		&controllers.FinderController{},
		// 	),
		// ),
		beego.NSNamespace("/Api/Sns",
			beego.NSInclude(
				&controllers.FriendCircleController{},
			),
		),
		beego.NSNamespace("/Api/Favor",
			beego.NSInclude(
				&controllers.FavorController{},
			),
		),
		beego.NSNamespace("/Api/Chatroom",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),
		beego.NSNamespace("/Api/Label",
			beego.NSInclude(
				&controllers.LabelController{},
			),
		),
		beego.NSNamespace("/Api/User",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		// beego.NSNamespace("/SayHello",
		// 	beego.NSInclude(
		// 		&controllers.SayHelloController{},
		// 	),
		// ),
		beego.NSNamespace("/Api/Common",
			beego.NSInclude(
				&controllers.ToolsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
