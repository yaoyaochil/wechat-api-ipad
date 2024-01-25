package Msg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"

	"github.com/golang/protobuf/proto"
)

type MassSendRequestParam struct {
	Wxid    string
	ToWxids string
	Content string
	Type    uint32
}

func MassSend(Data MassSendRequestParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	lists := strings.Split(Data.ToWxids, ";")
	md := md5.Sum([]byte(Data.Content))
	md5 := hex.EncodeToString(md[:])
	clientid := fmt.Sprintf("%v_%v", Data.Wxid, time.Now().Unix())
	buffer := []byte(Data.Content)

	//消息组包
	MsgRequest := &mm.MassSendRequest{
		ToList:      proto.String(Data.ToWxids),
		ToListCount: proto.Uint32(uint32(len(lists))),
		ToListMd5:   proto.String(md5),
		MsgType:     proto.Uint32(Data.Type),
		ClientID:    proto.String(clientid),
		DataBuffer: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(buffer))),
			Buffer: buffer,
		},
	}

	//序列化
	reqdata, _ := proto.Marshal(MsgRequest)

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:            D.Mmtlsip,
		Cgiurl:        "/cgi-bin/micromsg-bin/masssend",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              193,
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
	NewSendMsgRespone := mm.MassSendResponse{}
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
