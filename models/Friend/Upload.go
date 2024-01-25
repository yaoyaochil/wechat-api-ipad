package Friend

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	//"strings"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type UploadParam struct {
	Wxid            string
	PhoneNumberList []string
	Mobile          string
}

func UploadMContact(Data UploadParam, Opcode int32) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	if Data.PhoneNumberList == nil || len(Data.PhoneNumberList) == 0 {
		return wxCilent.ResponseResult{
			Code:    -9,
			Success: false,
			Message: "PhoneNumberList 手机号必填",
			Data:    nil,
		}
	}

	var PhoneNoList []*mm.SKBuiltinStringT

	for _, v := range Data.PhoneNumberList {
		PhoneNoList = append(PhoneNoList, &mm.SKBuiltinStringT{
			String_: proto.String(v),
		})
	}

	req := &mm.UploadMContactRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		UserName:       proto.String(Data.Wxid),
		Opcode:         proto.Int32(Opcode),
		Mobile:         proto.String(Data.Mobile),
		MobileListSize: proto.Int32(int32(len(PhoneNoList))),
		MobileList:     PhoneNoList,
		EmailListSize:  proto.Int32(0),
		EmailList:      nil,
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
		Cgiurl:        "/cgi-bin/micromsg-bin/uploadmcontact",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              133,
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
	Response := mm.UploadMContactResponse{}
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
