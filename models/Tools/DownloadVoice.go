package Tools

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

func DownloadVoice(Data DownloadVoiceParam) wxCilent.ResponseResult {
	var err error
	var protobufdata []byte
	var errtype int64

	var Response mm.DownloadVoiceResponse

	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	I := 0

	Startpos := 0
	datalen := 50000

	Databuff := make([]byte, Data.Length + 1000)
	var VoiceLength int32

	NewMsgId, _ :=  strconv.ParseInt(Data.NewMsgId, 10, 64)
	MasterBufId, _ := strconv.ParseInt(Data.Bufid, 10, 64)

	for   {
		Startpos = I * datalen
		count := 0
		if Data.Length - Startpos > datalen {
			count = Data.Length
		} else {
			count = Data.Length - Startpos
		}
		if count < 0 {
			break
		}

		req := &mm.DownloadVoiceRequest{
			MsgId:                proto.Int32(0),
			Offset:               proto.Int32(int32(Startpos)),
			Length:               proto.Int32(int32(count)),
			BaseRequest:          &mm.BaseRequest{
				SessionKey:    D.Sessionkey,
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
				DeviceType:    wxCilent.DeviceType_byte,
				Scene:         proto.Uint32(0),
			},
			NewMsgId:             proto.Int64(NewMsgId),
			ChatRoomName:         proto.String(Data.FromUserName),
			MasterBufId:          proto.Int64(MasterBufId),
		}

		//序列化
		reqdata, _ := proto.Marshal(req)

		//发包
		protobufdata, _, errtype, err = comm.SendRequest(comm.SendPostData{
			Ip:            D.Mmtlsip,
			Cgiurl:        "/cgi-bin/micromsg-bin/downloadvoice",
			Proxy:         D.Proxy,
			Encryption:    5,
			TwelveEncData: wxCilent.PackSpecialCgiData{},
			PackData: wxCilent.PackData{
				Reqdata:          reqdata,
				Cgi:              128,
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
			break
		}

		//解包
		err = proto.Unmarshal(protobufdata, &Response)
		if err != nil || *Response.BaseResponse.Ret != 0 {
			break
		}

		DataStream := bytes.NewBuffer(Response.Data.Buffer)
		_, _ = DataStream.Read(Databuff)
		VoiceLength = *Response.VoiceLength

		I++
	}

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    Response,
		}
	}

	if *Response.BaseResponse.Ret != 0 {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "异常：语音下载失败",
			Data:    Response,
		}
	}

	return wxCilent.ResponseResult{
		Code:    0,
		Success: false,
		Message: "成功",
		Data:   DownloadVoiceData{
			Base64:      Databuff,
			VoiceLength: VoiceLength,
		},
	}

}
