package Friend

import (
	"fmt"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"

	"github.com/golang/protobuf/proto"
)

type SendRequestParam struct {
	Wxid          string
	Opcode        int32
	V1            string
	V4            string
	Scene         int
	VerifyContent string
}

func SendRequest(Data SendRequestParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	VerifyUserList := make([]*mm.VerifyUser, 0)
	verifyTicket := ""
	if Data.Opcode == 3 {
		verifyTicket = Data.V4
	}

	VerifyUserList = append(VerifyUserList, &mm.VerifyUser{
		Value:               proto.String(Data.V1),
		VerifyUserTicket:    proto.String(verifyTicket),
		AntispamTicket:      proto.String(Data.V4),
		FriendFlag:          proto.Uint32(0),
		ChatRoomUserName:    proto.String(""),
		SourceUserName:      proto.String(""),
		SourceNickName:      proto.String(""),
		ScanQrcodeFromScene: proto.Uint32(0),
		ReportInfo:          proto.String(""),
		OuterUrl:            proto.String(""),
		SubScene:            proto.Int32(0),
	})

	ccData := &mm.CryptoData{
		Version:     []byte("00000003"),
		Type:        proto.Uint32(1),
		EncryptData: wxCilent.GetNewSpamData(D.Deviceid_str, D.DeviceName),
		Timestamp:   proto.Uint32(uint32(time.Now().Unix())),
		Unknown5:    proto.Uint32(5),
		Unknown6:    proto.Uint32(0),
	}
	ccDataseq, _ := proto.Marshal(ccData)
	WCExtInfo := &mm.WCExtInfo{
		CcData: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(ccDataseq))),
			Buffer: ccDataseq,
		},
	}

	WCExtInfoseq, _ := proto.Marshal(WCExtInfo)

	req := &mm.VerifyUserRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		Opcode:             proto.Int32(Data.Opcode),
		VerifyUserListSize: proto.Uint32(1),
		VerifyUserList:     VerifyUserList,
		VerifyContent:      proto.String(Data.VerifyContent),
		SceneList:          []byte{byte(Data.Scene)},
		SceneListCount:     proto.Uint32(1),
		ExtSpamInfo: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(WCExtInfoseq))),
			Buffer: WCExtInfoseq,
		},
		NeedConfirm: proto.Uint32(1),
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
		Cgiurl:        "/cgi-bin/micromsg-bin/verifyuser",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              137,
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
	Response := mm.VerifyUserResponse{}
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
