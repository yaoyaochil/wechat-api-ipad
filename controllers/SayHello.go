package controllers

import (
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// wxCilent "wechatwebapi/Cilent"
	// "wechatwebapi/models/SayHello"
)

// 打招呼模块
type SayHelloController struct {
	beego.Controller
}

// @Summary 模式-扫码
// @Param	body			body	SayHello.Model1Param	 true		"注意,请先执行1再执行2"
// @Failure 200
// @router /Modelv1 [post]
// func (c *SayHelloController) ModelV1() {
// 	var Data SayHello.Model1Param
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
// 	if err != nil {
// 		Result := wxCilent.ResponseResult{
// 			Code:    -8,
// 			Success: false,
// 			Message: fmt.Sprintf("系统异常：%v", err.Error()),
// 			Data:    nil,
// 		}
// 		c.Data["json"] = &Result
// 		c.ServeJSON()
// 		return
// 	}
// 	WXDATA := SayHello.Model1(Data)
// 	c.Data["json"] = &WXDATA
// 	c.ServeJSON()
// }
