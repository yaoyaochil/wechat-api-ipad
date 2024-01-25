package controllers

import (
	"github.com/astaxie/beego"
	// "wechatwebapi/models/Finder"
)

// 视频号
type FinderController struct {
	beego.Controller
}

// @Summary 用户中心
// @Param	wxid		query 	string	true		"请输登陆后的wxid"
// @Success 200
// @router /UserPrepare [post]
// func (c *FinderController) UserPrepare() {
// 	wxid := c.GetString("wxid")
// 	WXDATA := Finder.UserPrepare(wxid)
// 	c.Data["json"] = &WXDATA
// 	c.ServeJSON()
// }
