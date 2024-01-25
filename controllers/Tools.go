package controllers

import (
	"encoding/json"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Tools"

	"github.com/astaxie/beego"
)

// 工具
type ToolsController struct {
	beego.Controller
}

// @Summary 设置/删除代理IP
// @Param	body		body 	Tools.SetProxyParam   true	"删除代理ip时直接留空即可"
// @Success 200
// @router /setproxy [post]
func (c *ToolsController) SetProxy() {
	var ParamData Tools.SetProxyParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.SetProxy(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary GetA8Key
// @Param	body		body 	Tools.GetA8KeyParam   true	"OpCode = 2 Scene = 4 CodeType = 19 CodeVersion = 5 为默认参数,如有需求自行修改"
// @Success 200
// @router /GetA8Key [post]
func (c *ToolsController) GetA8Key() {
	var ParamData Tools.GetA8KeyParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.GetA8Key(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary MPGetA8Key
// @Param	body		body 	Tools.MPGetA8KeyParam   true	"scene 2=好友或群 3=历史阅读 4=二维码 7=公众号 30=扫码进群 opcode=2"
// @Success 200
// @router /MPGetA8Key [post]
func (c *ToolsController) MPGetA8Key() {
	var ParamData Tools.MPGetA8KeyParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.MPGetA8Key(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取个人或群二维码
// @Param	body		body 	Tools.GetQrRequestParam   true	"非获取群聊二维码时Chatroom字段留空 style:二维码样式 0-2"
// @Success 200
// @router /GetQrcode [post]
func (c *ToolsController) GetQrcode() {
	var ParamData Tools.GetQrRequestParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.GetQrcode(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 绑定手机
// @Param	body		body 	Tools.BindMobileRequestParam   true	"Opcode: 1:获取验证码 2:绑定手机号 3:解绑手机号"
// @Success 200
// @router /BindMobile [post]
func (c *ToolsController) BindMobile() {
	var ParamData Tools.BindMobileRequestParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.BindMobile(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 语音下载
// @Param	body		body 	Tools.DownloadVoiceParam    true	"同步消息获取参数"
// @Success 200
// @router /DownloadVoice [post]
func (c *ToolsController) DownloadVoice() {
	var ParamData Tools.DownloadVoiceParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.DownloadVoice(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取安全设备
// @Param	wxid		query 	string    true	"输入登录成功后的wxid"
// @Success 200
// @router /GetSafetyInfo [post]
func (c *ToolsController) GetSafetyInfo() {
	wxid := c.GetString("wxid")
	WXDATA := Tools.GetSafetyInfo(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 删除安全设备
// @Param	body		body 	Tools.DelSafeDeviceParam    true	"输入登录成功后的wxid"
// @Success 200
// @router /DelSafetyInfo [post]
func (c *ToolsController) DelSafetyInfo() {
	var ParamData Tools.DelSafeDeviceParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
	if err != nil {
		Result := wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}

	WXDATA := Tools.DelSafetyInfo(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
