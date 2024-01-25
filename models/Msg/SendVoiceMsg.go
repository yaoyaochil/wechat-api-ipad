package Msg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"

	"github.com/golang/protobuf/proto"
)

type SendVoiceMessageParam struct {
	Wxid      string
	ToWxid    string
	Base64    string
	VoiceTime int32
	Type      int32
}

func SendVoiceMsg(Data SendVoiceMessageParam) wxCilent.ResponseResult {
	var err error
	var protobufdata []byte
	var errtype int64

	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	VoiceData := strings.Split(Data.Base64, ",")

	var VoiceBase64 []byte

	if len(VoiceData) > 1 {
		VoiceBase64, _ = base64.StdEncoding.DecodeString(VoiceData[1])
	} else {
		VoiceBase64, _ = base64.StdEncoding.DecodeString(Data.Base64)
	}

	VoiceStream := bytes.NewBuffer(VoiceBase64)

	Startpos := 0
	datalen := 65000
	datatotalength := VoiceStream.Len()

	ClientImgId := fmt.Sprintf("%v_%v", Data.Wxid, time.Now().Unix())

	I := 0

	for {
		Startpos = I * datalen
		count := 0
		if datatotalength-Startpos > datalen {
			count = datalen
		} else {
			count = datatotalength - Startpos
		}
		if count < 0 {
			break
		}

		Databuff := make([]byte, count)
		_, _ = VoiceStream.Read(Databuff)

		req := &mm.UploadVoiceRequest{
			FromUserName: proto.String(Data.Wxid),
			ToUserName:   proto.String(Data.ToWxid),
			Offset:       proto.Uint32(uint32(Startpos)),
			Length:       proto.Int32(int32(datatotalength)),
			ClientMsgId:  proto.String(ClientImgId),
			MsgId:        proto.Uint32(0),
			VoiceLength:  proto.Int32(Data.VoiceTime),
			Data: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(Databuff))),
				Buffer: Databuff,
			},
			EndFlag: proto.Uint32(1),
			BaseRequest: &mm.BaseRequest{
				SessionKey:    D.Sessionkey,
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
				DeviceType:    wxCilent.DeviceType_byte,
				Scene:         proto.Uint32(0),
			},
			CancelFlag:  proto.Uint32(0),
			Msgsource:   proto.String(""),
			VoiceFormat: proto.Int32(Data.Type),
			ForwardFlag: proto.Uint32(0),
			NewMsgId:    proto.Uint64(0),
			Offst:       proto.Uint32(0),
		}

		//序列化
		reqdata, _ := proto.Marshal(req)

		//发包
		protobufdata, _, errtype, err = comm.SendRequest(comm.SendPostData{
			Ip:            D.Mmtlsip,
			Cgiurl:        "/cgi-bin/micromsg-bin/uploadvoice",
			Proxy:         D.Proxy,
			Encryption:    5,
			TwelveEncData: wxCilent.PackSpecialCgiData{},
			PackData: wxCilent.PackData{
				Reqdata:          reqdata,
				Cgi:              127,
				Uin:              D.Uin,
				Cookie:           D.Cooike,
				Sessionkey:       D.Sessionkey,
				EncryptType:      5,
				Loginecdhkey:     D.Loginecdhkey,
				Clientsessionkey: D.Clientsessionkey,
				UseCompress:      true,
			},
		}, D.MmtlsKey)

		if err != nil {
			break
		}

		I++
	}

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	//解包
	Response := mm.UploadVoiceResponse{}
	err = proto.Unmarshal(protobufdata, &Response)
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
		Success: false,
		Message: "成功",
		Data:    Response,
	}

}
