package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/FriendCircle"
)

// 朋友圈
type FriendCircleController struct {
	beego.Controller
}

// @Summary 同步朋友圈
// @Param	body		body 	FriendCircle.MmSnsSyncParam   true	"Synckey可留空"
// @Success 200
// @router /SnsSync [post]
func (c *FriendCircleController) SnsSync() {
	var ParamData FriendCircle.MmSnsSyncParam
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

	WXDATA := FriendCircle.MmSnsSync(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 朋友圈首页列表
// @Param	body			body	FriendCircle.GetListParam	true		"打开首页时：Fristpagemd5留空,Maxid填写0"
// @Failure 200
// @router /SnsTimeline [post]
func (c *FriendCircleController) SnsTimeline() {
	var Data FriendCircle.GetListParam
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
	WXDATA := FriendCircle.GetList(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取特定人朋友圈
// @Param	body			body	FriendCircle.GetDetailparameter	true		"打开首页时：Fristpagemd5留空,Maxid填写0"
// @Failure 200
// @router /SnsUserPage [post]
func (c *FriendCircleController) SnsUserPage() {
	var Data FriendCircle.GetDetailparameter
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
	WXDATA := FriendCircle.GetDetail(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 获取特定一条朋友圈详情内容
// @Param	body			body	FriendCircle.GetIdDetailParam	true		"Id为对象朋友圈内容的id"
// @Failure 200
// @router /SnsObjectDetail [post]
func (c *FriendCircleController) SnsObjectDetail() {
	var Data FriendCircle.GetIdDetailParam
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
	WXDATA := FriendCircle.GetIdDetail(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 发布朋友圈
// @Param	body			body	FriendCircle.Messagearameter	true		"请自行构造xml内容"
// @Failure 200
// @router /SnsPost [post]
func (c *FriendCircleController) SnsPost() {
	var Data FriendCircle.Messagearameter
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
	WXDATA := FriendCircle.Messages(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 朋友圈操作
// @Param	body			body	FriendCircle.OperationParam	true		"type：1删除朋友圈2设为隐私3设为公开4删除评论5取消点赞"
// @Failure 200
// @router /SnsObjectTop [post]
func (c *FriendCircleController) SnsObjectTop() {
	var Data FriendCircle.OperationParam
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
	WXDATA := FriendCircle.Operation(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 朋友圈点赞/评论
// @Param	body			body	FriendCircle.CommentParam	true		"type：1点赞 2：文本 3:消息 4：with 5陌生人点赞 replyCommnetId：回复评论Id"
// @Failure 200
// @router /SnsComment [post]
func (c *FriendCircleController) SnsComment() {
	var Data FriendCircle.CommentParam
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
	WXDATA := FriendCircle.Comment(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 朋友圈上传
// @Param	body			body	FriendCircle.SnsUploadParam 	true		"支持图片和视频"
// @Failure 200
// @router /SnsUpload [post]
func (c *FriendCircleController) SnsUpload() {
	var Data FriendCircle.SnsUploadParam
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
	WXDATA := FriendCircle.SnsUpload(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 朋友圈权限设置
// @Param	body	body	FriendCircle.PrivacySettingsParam	 true		"调用参数请联系"
// @Success 200
// @router /SetSnsViewTime [post]
func (c *FriendCircleController) SetSnsViewTime() {
	var Data FriendCircle.PrivacySettingsParam
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
	WXDATA := FriendCircle.PrivacySettings(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
