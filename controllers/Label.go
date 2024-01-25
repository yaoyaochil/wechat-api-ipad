package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/models/Label"
)

// 标签
type LabelController struct {
	beego.Controller
}

// @Summary 获取标签列表
// @Param	wxid			query 	string	true		""
// @Failure 200
// @router /GetContactLabelList [post]
func (c *LabelController) GetContactLabelList() {
	wxid := c.GetString("wxid")
	WXDATA := Label.GetList(wxid)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 添加标签
// @Param	body			body	Label.AddParam	true		""
// @Failure 200
// @router /AddContactLabel [post]
func (c *LabelController) AddContactLabel() {
	var Data Label.AddParam
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
	WXDATA := Label.Add(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 更新标签列表
// @Param	body			body	Label.UpdateListParam	true		"ToWxids: 要更新label的对象wxid"
// @Failure 200
// @router /UpdateContactLabelList [post]
func (c *LabelController) UpdateContactLabelList() {
	var Data Label.UpdateListParam
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
	WXDATA := Label.UpdateList(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 修改标签
// @Param	body			body	Label.UpdateNameParam	true		""
// @Failure 200
// @router /UpdateContactLabel [post]
func (c *LabelController) UpdateContactLabel() {
	var Data Label.UpdateNameParam
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
	WXDATA := Label.UpdateName(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 删除标签
// @Param	body			body	Label.DeleteParam	true		""
// @Failure 200
// @router /DeleteContactLabel [post]
func (c *LabelController) DeleteContactLabel() {
	var Data Label.DeleteParam
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
	WXDATA := Label.Delete(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
