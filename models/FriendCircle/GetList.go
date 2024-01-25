package FriendCircle

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type GetListParam struct {
	Wxid         string
	Maxid        uint64
	Fristpagemd5 string
}

func GetList(Data GetListParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	req := &mm.SnsTimeLineRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(369558056),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		FirstPageMd5:    proto.String(Data.Fristpagemd5),
		MaxId:           proto.Uint64(Data.Maxid),
		MinFilterId:     proto.Uint64(0),
		LastRequestTime: proto.Uint32(0),
		ClientLatestId:  proto.Uint64(0),
		NetworkType:     proto.Int32(1),
	}

	reqdata, err := proto.Marshal(req)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:            D.Mmtlsip,
		Cgiurl:        "/cgi-bin/micromsg-bin/mmsnstimeline",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              683,
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
		return wxCilent.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	//解包
	Response := mm.SnsTimeLineResponse{}
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
		Success: true,
		Message: "成功",
		Data:    Response,
	}
}
