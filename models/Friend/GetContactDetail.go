package Friend


import (
	"fmt"
	"github.com/golang/protobuf/proto"
	//"strings"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type GetContactDetailparameter struct {
	Wxid     string
	Towxids    []string
	ChatRoom string
}

func GetContact(Data GetContactDetailparameter) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	Towxds := make([]*mm.SKBuiltinStringT, len(Data.Towxids))

	if len(Data.Towxids) >= 1 {
		for i, v := range Data.Towxids {
			Towxds[i] = &mm.SKBuiltinStringT{
				String_: proto.String(v),
			}
		}
	}

	ChatRoom := &mm.SKBuiltinStringT{}
	ChatRoomCount := uint32(1)

	if Data.ChatRoom != "" {
		ChatRoomCount = 1
		ChatRoom = &mm.SKBuiltinStringT{
			String_: proto.String(Data.ChatRoom),
		}
	} else {
		ChatRoom = nil
		ChatRoomCount = uint32(0)
	}

	req := &mm.GetContactRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		UserCount:         proto.Int32(int32(len(Towxds))),
		UserNameList:      Towxds,
		FromChatRoomCount: proto.Int32(int32(ChatRoomCount)),
		FromChatRoom:      ChatRoom,
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
		Cgiurl:        "/cgi-bin/micromsg-bin/getcontact",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              182,
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
	GetContactResponse := mm.GetContactResponse{}
	err = proto.Unmarshal(protobufdata, &GetContactResponse)

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
		Data:    GetContactResponse,
	}
}
