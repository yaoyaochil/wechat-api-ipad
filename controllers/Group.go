package controllers

import (
	"encoding/json"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Group"
	"wechatwebapi/models/Tools"

	"github.com/astaxie/beego"
)

// 群组
type GroupController struct {
	beego.Controller
}

// @Summary 创建群聊
// @Param	body			body	Group.CreateChatRoomParam	true		"ToWxids 多个微信ID用,隔开 至少三个好友微信ID以上"
// @Failure 200
// @router /CreateChatRoom [post]
func (c *GroupController) CreateChatRoom() {
	var Data Group.CreateChatRoomParam
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
	WXDATA := Group.CreateChatRoom(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 增加群成员(40人以内)
// @Param	body			body	Group.AddChatRoomParam	true		"ToWxids 多个微信ID用,隔开 QID 群ID"
// @Failure 200
// @router /AddChatRoomMember [post]
func (c *GroupController) AddChatRoomMember() {
	var Data Group.AddChatRoomParam
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
	WXDATA := Group.AddChatRoomMember(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 邀请群成员(40人以上)
// @Param	body			body	Group.AddChatRoomParam	true		"ToWxids 多个微信ID用,隔开 QID: 群ID"
// @Failure 200
// @router /InviteChatRoomMember [post]
func (c *GroupController) InviteChatRoomMember() {
	var Data Group.AddChatRoomParam
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
	WXDATA := Group.InviteChatRoomMember(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 删除群成员
// @Param	body			body	Group.AddChatRoomParam	true		"ToWxids:多个微信ID用,隔开 QID: 群ID"
// @Failure 200
// @router /DelChatRoomMember [post]
func (c *GroupController) DelChatRoomMember() {
	var Data Group.AddChatRoomParam
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
	WXDATA := Group.DelChatRoomMember(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 扫码进群
// @Param	body			body	Group.ScanIntoGroupParam	true		"url"
// @Failure 200
// @router /ScanIntoGroup [post]
func (c *GroupController) ScanIntoGroup() {
	var Data Group.ScanIntoGroupParam
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
	WXDATA := Group.ScanIntoGroup(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 退出群聊
// @Param	body			body	Group.QuitGroupParam 	true		"QID = 群ID"
// @Failure 200
// @router /QuitChatroom [post]
func (c *GroupController) QuitChatroom() {
	var Data Group.QuitGroupParam
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
	WXDATA := Group.Quit(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取群详情(不带公告内容)
// @Param	body			body	Group.GetChatRoomParam 	true		"UserNameList == 群ID,多个查询请用,隔开"
// @Failure 200
// @router /GetChatRoomInfo [post]
func (c *GroupController) GetChatRoomInfo() {
	var Data Group.GetChatRoomParam
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
	WXDATA := Tools.GetContact(Tools.GetContactParam{
		Wxid:         Data.Wxid,
		UserNameList: Data.QID,
	})
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取群信息(带公告内容)
// @Param	body			body	Group.GetChatRoomParam 	true		"QID == 群ID"
// @Failure 200
// @router /GetChatRoomInfoDetail [post]
func (c *GroupController) GetChatRoomInfoDetail() {
	var Data Group.GetChatRoomParam
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
	WXDATA := Group.GetChatRoomInfoDetail(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取群成员详情
// @Param	body			body	Group.GetChatRoomParam 	true		"QID = 群ID"
// @Failure 200
// @router /GetChatRoomMemberDetail [post]
func (c *GroupController) GetChatRoomMemberDetail() {
	var Data Group.GetChatRoomParam
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
	WXDATA := Group.GetChatRoomMemberDetail(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 设置群备注(仅自己可见)
// @Param	body			body	Group.OperateChatRoomInfoParam 	true		"QID = 群ID"
// @Failure 200
// @router /SetChatRoomRemarks [post]
func (c *GroupController) SetChatRoomRemarks() {
	var Data Group.OperateChatRoomInfoParam
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
	WXDATA := Group.SetChatRoomRemarks(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 保存到通讯录
// @Param	body			body	Group.MoveContractListParam 	true		"Val = 1添加 0移除"
// @Failure 200
// @router /MoveToContract [post]
func (c *GroupController) MoveToContract() {
	var Data Group.MoveContractListParam
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
	WXDATA := Group.MoveToContract(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 设置群名称
// @Param	body			body	Group.OperateChatRoomInfoParam 	true		"Content = 名称"
// @Failure 200
// @router /SetChatRoomName [post]
func (c *GroupController) SetChatRoomName() {
	var Data Group.OperateChatRoomInfoParam
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
	WXDATA := Group.SetChatRoomName(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 群管理操作(添加、删除、转让)
// @Param	body			body	Group.OperateChatRoomAdminParam 	true		"ToWxids = 多个wxid用,隔开(仅限于添加/删除管理员) Val = 1添加 2删除 3转让"
// @Failure 200
// @router /OperateChatRoomAdmin [post]
func (c *GroupController) OperateChatRoomAdmin() {
	var Data Group.OperateChatRoomAdminParam
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
	WXDATA := Group.OperateChatRoomAdmin(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 设置群公告
// @Param	body			body	Group.OperateChatRoomInfoParam 	true		"Content = 公告内容"
// @Failure 200
// @router /SetChatRoomAnnouncement [post]
func (c *GroupController) SetChatRoomAnnouncement() {
	var Data Group.OperateChatRoomInfoParam
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
	WXDATA := Group.SetChatRoomAnnouncement(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
