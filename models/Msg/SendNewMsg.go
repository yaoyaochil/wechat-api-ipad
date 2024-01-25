package Msg

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type SendNewMsgParam struct {
	Wxid    string
	ToWxid  string
	Content string
	Type    int64
}

func SendNewMsg(Data SendNewMsgParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	//消息组包
	MsgRequest := &mm.NewSendMsgRequest{
		Cnt: proto.Int32(1),
		Info: &mm.ChatInfo{
			Toid: &mm.SKBuiltinStringT{
				String_: proto.String(Data.ToWxid),
			},
			Content:     proto.String(Data.Content),
			Type:        proto.Int64(Data.Type),
			Utc:         proto.Int64(time.Now().Unix()),
			ClientMsgId: proto.Uint64(uint64(time.Now().Unix() + 567073593)),
			MsgSource:   nil,
		},
	}

	//序列化
	reqdata, _ := proto.Marshal(MsgRequest)

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:            D.Mmtlsip,
		Cgiurl:        "/cgi-bin/micromsg-bin/newsendmsg",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              522,
			Uin:              D.Uin,
			Cookie:           D.Cooike,
			Sessionkey:       D.Sessionkey,
			EncryptType:      5,
			Loginecdhkey:     D.Loginecdhkey,
			Clientsessionkey: D.Clientsessionkey,
			UseCompress:      false,
		},
	}, D.MmtlsKey)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	//解包
	NewSendMsgRespone := mm.NewSendMsgRespone{}
	err = proto.Unmarshal(protobufdata, &NewSendMsgRespone)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	return wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    NewSendMsgRespone,
	}

}
