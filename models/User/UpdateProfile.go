package User

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type UpdateProfileParam struct {
	Wxid      string
	NickName  string
	Sex       int32
	Country   string
	Province  string
	City      string
	Signature string
}

func UpdateProfile(Data UpdateProfileParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	info := GetContractProfile(Data.Wxid)

	MM, _ := json.Marshal(info.Data)

	var ModUserInfo mm.GetProfileResponse

	_ = json.Unmarshal(MM, &ModUserInfo)

	userInfo := &mm.ModUserInfo{}

	if Data.NickName != "" {
		userInfo.NickName = &mm.SKBuiltinStringT{
			String_: proto.String(Data.NickName),
		}
	}

	if Data.City != "" {
		userInfo.City = proto.String(Data.City)
	}

	if Data.Country != "" {
		userInfo.Country = proto.String(Data.Country)
	}

	if Data.Province != "" {
		userInfo.Province = proto.String(Data.Province)
	}

	if Data.Signature != "" {
		userInfo.Signature = proto.String(Data.Signature)
	}

	userInfo.UserName = &mm.SKBuiltinStringT{
		String_: proto.String(Data.Wxid),
	}

	userInfo.BitFlag = proto.Uint32(*ModUserInfo.UserInfo.BitFlag)
	userInfo.Status = proto.Uint32(*ModUserInfo.UserInfo.Status)
	userInfo.PluginFlag = proto.Uint32(*ModUserInfo.UserInfo.PluginFlag)
	userInfo.BindMobile = ModUserInfo.UserInfo.BindMobile
	userInfo.BindUin = ModUserInfo.UserInfo.BindUin
	userInfo.BindEmail = ModUserInfo.UserInfo.BindEmail
	userInfo.ImgLen = ModUserInfo.UserInfo.ImgLen

	userInfo.Sex = proto.Int32(Data.Sex)

	//序列化
	userInfoSerialize, _ := proto.Marshal(userInfo)

	var CmdItem []*mm.CmdItem

	CmdItem = append(CmdItem, &mm.CmdItem{
		CmdId: proto.Int32(1),
		CmdBuf: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(userInfoSerialize))),
			Buffer: userInfoSerialize,
		},
	})

	req := &mm.OpLogRequest{
		Cmd: &mm.CmdList{
			Count: proto.Uint32(uint32(len(CmdItem))),
			List:  CmdItem,
		},
	}

	//序列化
	reqdata, _ := proto.Marshal(req)

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:            D.Mmtlsip,
		Cgiurl:        "/cgi-bin/micromsg-bin/oplog",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              681,
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
	Response := mm.OplogResponse{}
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
