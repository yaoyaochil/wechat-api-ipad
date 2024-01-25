package Msg

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

func SendAppMsg(Data SendAppMsgParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	time := time.Now().Unix()

	req := &mm.SendAppMsgRequest{
		BaseRequest:          &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		Msg:                  &mm.AppMsg{
			FromUserName:         proto.String(Data.Wxid),
			AppId:                proto.String(""),
			SdkVersion:           proto.Int32(0),
			ToUserName:           proto.String(Data.ToWxid),
			Type:                 proto.Int32(Data.Type),
			Content:              proto.String(Data.Content),
			CreateTime:           proto.Int64(time),
			ClientMsgId:          proto.String(fmt.Sprintf("%v_%v",Data.ToWxid,time)),
			Source:               proto.Int32(0),
			MsgSource:            proto.String("<msgsource><bizflag>0</bizflag></msgsource>"),
		},
		MsgForwardType:       proto.Int32(2),
		SendMsgTicket:        proto.String(""),
	}

	//序列化
	reqdata, _ := proto.Marshal(req)

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:            D.Mmtlsip,
		Cgiurl:        "/cgi-bin/micromsg-bin/sendappmsg",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              222,
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
	NewSendMsgRespone := mm.SendAppMsgResponse{}
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
