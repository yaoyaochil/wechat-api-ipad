package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/User"
)

// 微信号管理
type UserController struct {
	beego.Controller
}

// @Summary 取个人信息
// @Param	wxid		query 	string	true		"请输入登陆后的wxid"
// @Success 200
// @router /GetProfile [post]
func (c *UserController) GetProfile() {
	wxid := c.GetString("wxid")
	WXDATA := User.GetContractProfile(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 隐私设置
// @Param	body	body	User.PrivacySettingsParam	 true		"调用参数请咨询"
// @Success 200
// @router /VerifySwitch [post]
func (c *UserController) VerifySwitch() {
	var Data User.PrivacySettingsParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
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
	WXDATA := User.PrivacySettings(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 修改个人信息
// @Param	body	body	User.UpdateProfileParam	 true		"NickName=昵称  Sex=性别（0:位置 1:男 2：女） Country=国家 Province=省份 City=城市 Signature=个性签名"
// @Success 200
// @router /UpdateProfile [post]
func (c *UserController) UpdateProfile() {
	var Data User.UpdateProfileParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
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
	WXDATA := User.UpdateProfile(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 修改头像
// @Param	body	body	User.UploadHeadImageParam	 true		""
// @Success 200
// @router /UploadHeadImage [post]
func (c *UserController) UploadHeadImage() {
	var Data User.UploadHeadImageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
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
	WXDATA := User.UploadHeadImage(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 验证密码
// @Param	body	body	User.NewVerifyPasswdParam	 true		"调用SetPassword前调用该接口"
// @Success 200
// @router /VerifyPassword [post]
func (c *UserController) VerifyPassword() {
	var Data User.NewVerifyPasswdParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
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
	WXDATA := User.NewVerifyPasswd(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 修改密码
// @Param	body	body	User.NewSetPasswdParam	 true		"Ticket从VerifyPassword获取"
// @Success 200
// @router /SetPassword [post]
func (c *UserController) SetPassword() {
	var Data User.NewSetPasswdParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
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
	WXDATA := User.NewSetPasswd(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
