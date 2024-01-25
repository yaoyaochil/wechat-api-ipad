package FriendCircle

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type SnsUploadParam struct {
	Wxid   string
	Base64 string
}

func SnsUpload(Data SnsUploadParam) wxCilent.ResponseResult {
	var err error
	var protobufdata []byte
	var errtype int64
	var Bs64Data []byte

	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	Base64Data := strings.Split(Data.Base64, ",")

	if len(Base64Data) > 1 {
		Bs64Data, _ = base64.StdEncoding.DecodeString(Base64Data[1])
	} else {
		Bs64Data, _ = base64.StdEncoding.DecodeString(Data.Base64)
	}

	Stream := bytes.NewBuffer(Bs64Data)

	Bs64MD5 := wxCilent.GetFileMD5Hash(Bs64Data)

	Startpos := 0
	datalen := 50000
	datatotalength := Stream.Len()

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
		_, _ = Stream.Read(Databuff)

		req := &mm.SnsUploadRequest{
			BaseRequest: &mm.BaseRequest{
				SessionKey:    D.Sessionkey,
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
				DeviceType:    wxCilent.DeviceType_byte,
				Scene:         proto.Uint32(0),
			},
			Type:     proto.Uint32(2),
			StartPos: proto.Uint32(uint32(Startpos)),
			TotalLen: proto.Uint32(uint32(datatotalength)),
			Buffer: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(Databuff))),
				Buffer: Databuff,
			},
			ClientId: proto.String(ClientImgId),
			MD5:      proto.String(Bs64MD5),
		}

		//序列化
		reqdata, _ := proto.Marshal(req)

		//发包
		protobufdata, _, errtype, err = comm.SendRequest(comm.SendPostData{
			Ip:            D.Mmtlsip,
			Cgiurl:        "/cgi-bin/micromsg-bin/mmsnsupload",
			Proxy:         D.Proxy,
			Encryption:    5,
			TwelveEncData: wxCilent.PackSpecialCgiData{},
			PackData: wxCilent.PackData{
				Reqdata:          reqdata,
				Cgi:              207,
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
	Response := mm.SnsUploadResponse{}
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
