package controllers

import (
	"encoding/json"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Msg"

	"github.com/astaxie/beego"
)

// 消息
type MsgController struct {
	beego.Controller
}

// @Summary 同步消息
// @Param	body body Msg.SyncParam true "Scene填写0, Synckey留空"
// @Success 200
// @router /NewSync [post]
func (c *MsgController) NewSync() {
	var ParamData Msg.SyncParam
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

	WXDATA := Msg.Sync(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送App消息
// @Param	body		body 	Msg.SendAppMsgParam   true	"Type请根据场景设置, Content请自行构造"
// @Success 200
// @router /SendAppMsg [post]
func (c *MsgController) SendAppMsg() {
	var ParamData Msg.SendAppMsgParam
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

	WXDATA := Msg.SendAppMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送文本消息
// @Param	body		body 	Msg.SendNewMsgParam   true	"Type请填写1"
// @Success 200
// @router /SendMsg [post]
func (c *MsgController) SendMsg() {
	var ParamData Msg.SendNewMsgParam
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

	WXDATA := Msg.SendNewMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送图片
// @Param	body		body 	Msg.SendImageMsgParam   true	"base64为图片base64编码后数据"
// @Success 200
// @router /SendImg [post]
func (c *MsgController) SendImg() {
	var ParamData Msg.SendImageMsgParam
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

	WXDATA := Msg.SendImageMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送语音
// @Param	body		body 	Msg.SendVoiceMessageParam   true	"Type： AMR = 0, MP3 = 2, SILK = 4, SPEEX = 1, WAVE = 3 VoiceTime ：音频长度 1000为一秒"
// @Success 200
// @router /SendVoice [post]
func (c *MsgController) SendVoice() {
	var ParamData Msg.SendVoiceMessageParam
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

	WXDATA := Msg.SendVoiceMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送视频
// @Param	body		body 	Msg.SendVideoMsgParam   true	"Video:视频Base64 Image:封面Base64 VideoTime:视频长度"
// @Success 200
// @router /SendVideo [post]
func (c *MsgController) SendVideo() {
	var ParamData Msg.SendVideoMsgParam
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

	WXDATA := Msg.SendVideoMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 分享名片
// @Param	body		body 	Msg.ShareCardParam   true	"ToWxid=接收的微信ID Id=名片wxid NickName=名片昵称 Alias=名片别名 "
// @Success 200
// @router /ShareCard [post]
func (c *MsgController) ShareCard() {
	var ParamData Msg.ShareCardParam
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

	WXDATA := Msg.SendNewMsg(Msg.SendNewMsgParam{
		Wxid:    ParamData.Wxid,
		ToWxid:  ParamData.ToWxid,
		Content: fmt.Sprintf("<msg username=\"%v\" nickname=\"%v\" fullpy=\"\" shortpy=\"\" alias=\"%v\" imagestatus=\"3\" scene=\"17\" province=\"\" city=\"\" sign=\"\" sex=\"1\" certflag=\"0\" certinfo=\"\" brandIconUrl=\"\" brandHomeUrl=\"\" brandSubscriptConfigUrl=\"\" brandFlags=\"0\" regionCode=\"CN\" ></msg>", ParamData.Id, ParamData.NickName, ParamData.Alias),
		Type:    42,
	})
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 分享位置
// @Param	body		body 	Msg.ShareLocationParam   true	"x=经度 y=纬度 scale=地图缩放倍数，可填10，100实验"
// @Success 200
// @router /ShareLocation [post]
func (c *MsgController) ShareLocation() {
	var ParamData Msg.ShareLocationParam
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

	WXDATA := Msg.SendNewMsg(Msg.SendNewMsgParam{
		Wxid:    ParamData.Wxid,
		ToWxid:  ParamData.ToWxid,
		Content: fmt.Sprintf("<msg><location x=\"%v\" y=\"%v\" scale=\"%v\" label=\"%v\" poiname=\"%v\" maptype=\"roadmap\" infourl=\"\" fromusername=\"\" poiid=\"City\" /></msg>", ParamData.X, ParamData.Y, ParamData.Scale, ParamData.Title, ParamData.Poiname),
		Type:    48,
	})
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 撤回消息
// @Param	body		body 	Msg.RevokeMsgParam   true	"记录之前发送的消息的id与时间进行撤销"
// @Success 200
// @router /RevokeMsg [post]
func (c *MsgController) RevokeMsg() {
	var ParamData Msg.RevokeMsgParam
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

	WXDATA := Msg.RevokeMsg(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 群发助手
// @Param	body		body 	Msg.MassSendRequestParam   true	"ToWxids: 群发对象的wxid，用;分割组成一个字符串 Content:文本内容 Type: 1=文字消息"
// @Success 200
// @router /MassSend [post]
func (c *MsgController) MassSend() {
	var ParamData Msg.MassSendRequestParam
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

	WXDATA := Msg.MassSend(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
