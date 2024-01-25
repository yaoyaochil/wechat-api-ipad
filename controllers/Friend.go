package controllers

import (
	"encoding/json"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Friend"

	"github.com/astaxie/beego"
)

// 朋友
type FriendController struct {
	beego.Controller
}

// @Summary 搜索联系人
// @Param	body			body	Friend.SearchParam	 true		"默认填写FromScene=0,SearchScene=1"
// @Failure 200
// @router /SearchContact [post]
func (c *FriendController) SearchContact() {
	var Data Friend.SearchParam
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
	WXDATA := Friend.Search(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发送好友请求
// @Param	body			body	Friend.SendRequestParam	 true		"Scene:来源, VerifyContent: 招呼消息"
// @Failure 200
// @router /VerifyUser [post]
func (c *FriendController) VerifyUser() {
	var Data Friend.SendRequestParam
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
	WXDATA := Friend.SendRequest(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 上传通讯录
// @Param	body			body	Friend.UploadParam	 true		"PhoneNumberList:要上传的手机号 Mobile:自己的手机号"
// @Failure 200
// @router /UploadMContact [post]
func (c *FriendController) UploadMContact() {
	var Data Friend.UploadParam
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
	WXDATA := Friend.UploadMContact(Data, 1)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 移出通讯录
// @Param	body			body	Friend.UploadParam	 true		"PhoneNumberList:要删除的手机号 Mobile:自己的手机号"
// @Failure 200
// @router /DeleteFromContact [post]
func (c *FriendController) DeleteFromContact() {
	var Data Friend.UploadParam
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
	WXDATA := Friend.UploadMContact(Data, 2)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取通讯录好友
// @Param	wxid		query 	string	true		""
// @Failure 200
// @router /GetMFriend [post]
func (c *FriendController) GetMFriend() {
	wxid := c.GetString("wxid")
	WXDATA := Friend.GetMFriend(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取通讯录好友
// @Param	body			body	Friend.GetContactListparameter	 true		"CurrentWxcontactSeq和CurrentChatRoomContactSeq没有的情况下填写0"
// @Failure 200
// @router /GetContactList [post]
func (c *FriendController) GetContactList() {
	var Data Friend.GetContactListparameter
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
	WXDATA := Friend.GetContactList(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取通讯录好友详情
// @Param	body			body	Friend.GetContactDetailparameter	 true		"Towxids 要获取的微信(最多20个,建议15个), ChatRoom查询群里成员以外留空"
// @Failure 200
// @router /GetContact [post]
func (c *FriendController) GetContact() {
	var Data Friend.GetContactDetailparameter
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
	WXDATA := Friend.GetContact(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取好友是否删除拉黑自己
// @Param	body			body	Friend.GetContactStatusParameter	 true		"1:正常 0:删除 -1:拉黑"
// @Failure 200
// @router /GetContactStatus [post]
func (c *FriendController) GetContactStatus() {
	var Data Friend.GetContactStatusParameter
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
	WXDATA := Friend.GetContactStatus(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 设置好友备注
// @Param	body			body	Friend.SetRemarksParam	 true		""
// @Failure 200
// @router /SetContactRemarks [post]
func (c *FriendController) SetContactRemarks() {
	var Data Friend.SetRemarksParam
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
	WXDATA := Friend.SetRemarks(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 添加/移除黑名单
// @Param	body			body	Friend.BlacklistParam	 true		"Enable: 0 移出黑名单 1 移入黑名单"
// @Failure 200
// @router /SetContactBlacklist [post]
func (c *FriendController) SetContactBlacklist() {
	var Data Friend.BlacklistParam
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
	WXDATA := Friend.SetContactBlacklist(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 删除好友
// @Param	body			body	Friend.DelFriendParam	 true		"UserName: 要删除的好友wxid"
// @Failure 200
// @router /DelFriend [post]
func (c *FriendController) DelFriend() {
	var Data Friend.DelFriendParam
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
	WXDATA := Friend.DelFriend(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
