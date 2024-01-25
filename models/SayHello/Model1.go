package SayHello

import (
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/bts"
	"wechatwebapi/comm"
	"wechatwebapi/models/Friend"
)

type Model1Param struct {
	Wxid string
	Url  string
}

//扫码打招呼
func Model1(Data Model1Param) wxCilent.ResponseResult {
	_, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	if Data.Url == "" {
		return wxCilent.ResponseResult{
			Code:    -9,
			Success: false,
			Message: "Url是二维码的URL地址",
			Data:    nil,
		}
	}

	//搜索V1V2
	S := Friend.Search(Friend.SearchParam{
		Wxid:        Data.Wxid,
		ToUserName:  Data.Url,
		FromScene:   0,
		SearchScene: 1,
	})

	Search := bts.SearchContactResponse(S.Data)

	if *Search.AntispamTicket == "" {
		return wxCilent.ResponseResult{
			Code:    -9,
			Success: false,
			Message: "V2提取失败,请检查二维码URL是否正确或有效",
			Data:    nil,
		}
	}

	//开始提交验证
	Verify := Friend.SendRequest(Friend.SendRequestParam{
		Wxid:          Data.Wxid,
		V1:            *Search.UserName.String_,
		V2:            *Search.AntispamTicket,
		Scene:         30,
		VerifyContent: "",
	})

	return wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    Verify.Data,
	}

}
