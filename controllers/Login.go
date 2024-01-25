package controllers

import (
	"encoding/json"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Login"
	"github.com/astaxie/beego"
)

// 登陆
type LoginController struct {
	beego.Controller
}

// @Summary 获取登录二维码
// @Param	body		body 	Login.GetQRReq	true		"不使用代理留空"
// @Success 200
// @router /GetLoginQrcode [post]
func (c *LoginController) GetLoginQrcode() {
	/*//获取相关秘钥
	HybridEcdhInitServerPubKey, HybridEcdhPrivkey, HybridEcdhPubkey := wxCilent.HybridEcdhInit()
	httpclient := wxCilent.GenNewHttpClient()
	ret := httpclient.InitMmtlsShake(wxCilent.MMtls_ip)
	if !ret {

	}*/
	var GetQR Login.GetQRReq
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &GetQR)
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
	//生成设备ID
	DeviceID := wxCilent.CreateDeviceId(GetQR.DeviceID)
	if GetQR.DeviceName == "" {
		GetQR.DeviceName = "iPad"
	}

	WXDATA := Login.GetQRCODE(DeviceID, GetQR.DeviceName, GetQR.Proxy)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 检测二维码
// @Param	uuid		query 	string	true		"请输入取码时返回的UUID"
// @Success 200
// @router /CheckLoginQrcode [post]
func (c *LoginController) CheckLoginQrcode() {
	uuid := c.GetString("uuid")
	WXDATA := Login.CheckUuid(uuid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 二次登陆
// @Param	wxid			query 	string	true		"请输入登陆成功的wxid"
// @Failure 200
// @router /AutoAuth [post]
func (c *LoginController) AutoAuth() {
	wxid := c.GetString("wxid")
	WXDATA := Login.Secautoauth(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 账号密码登陆
// @Param	body			body 	Login.ManualAuthReq	true		"不使用代理留空"
// @Failure 200
// @router /Manualauth [post]
func (c *LoginController) Manualauth() {
	var reqdata Login.ManualAuthReq
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &reqdata)

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

	WXDATA := Login.Data62Login(reqdata)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 心跳
// @Param	wxid			query 	string	true		"请输入登陆成功的wxid"
// @Success 200
// @router /HeartBeat [post]
func (c *LoginController) HeartBeat() {
	wxid := c.GetString("wxid")
	WXDATA := Login.HeartBeat(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 初次登录初始化
// @Param	wxid			query 	string	true		"请输入登陆成功的wxid"
// @Param	MaxSynckey		query 	string	false		"二次同步需要带入"
// @Param	CurrentSynckey	query 	string	false		"二次同步需要带入"
// @Success 200
// @router /Newinit [post]
func (c *LoginController) Newinit() {
	wxid := c.GetString("wxid")
	MaxSynckey := c.GetString("MaxSynckey")
	CurrentSynckey := c.GetString("CurrentSynckey")
	WXDATA := Login.Newinit(wxid, MaxSynckey, CurrentSynckey)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 唤醒登陆(只限扫码登录)
// @Param	wxid		query 	string	true		"请输入登陆成功的wxid"
// @Success 200
// @router /Awaken [post]
func (c *LoginController) Awaken() {
	wxid := c.GetString("wxid")
	WXDATA := Login.AwakenLogin(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取登陆缓存信息
// @Param	wxid		query 	string	true		"请输入登陆成功的wxid"
// @Success 200
// @router /GetCacheInfo [post]
func (c *LoginController) GetCacheInfo() {
	wxid := c.GetString("wxid")
	WXDATA := Login.CacheInfo(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取62数据
// @Param	wxid		query 	string	true		"请输入登陆成功的wxid"
// @Success 200
// @router /Get62Data [post]
func (c *LoginController) Get62Data() {
	wxid := c.GetString("wxid")
	Data62 := Login.Get62Data(wxid)
	Result := wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    Data62,
	}
	c.Data["json"] = &Result
	c.ServeJSON()
	return
}

// @Summary 退出登录
// @Param	wxid			query 	string	true		"请输入登陆成功的wxid"
// @Success 200
// @router /LogOut [post]
func (c *LoginController) LogOut() {
	wxid := c.GetString("wxid")
	WXDATA := Login.LogOut(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
