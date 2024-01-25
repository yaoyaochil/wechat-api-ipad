package FriendCircle

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
	"wechatwebapi/models/Msg"
)

type MmSnsSyncParam struct {
	Wxid    string
	Synckey string
}

func MmSnsSync(Data MmSnsSyncParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	var Synckey mm.SKBuiltinBufferT

	if Data.Synckey != "" {
		key, _ := base64.StdEncoding.DecodeString(Data.Synckey)
		Synckey = mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(key))),
			Buffer: key,
		}
	}

	req := &mm.SnsSyncRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    []byte{},
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		Selector: proto.Uint32(256),
		KeyBuf:   &Synckey,
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
		Host:          D.MmtlsHost,
		Cgiurl:        "/cgi-bin/micromsg-bin/mmsnssync",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              214,
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
	Response := mm.SnsSyncResponse{}
	err = proto.Unmarshal(protobufdata, &Response)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	UnknownCmdId := ""

	var ModUserInfos []mm.ModUserInfo
	var ModContacts []mm.ModContact
	var DelContacts []mm.DelContact
	var ModUserImgs []mm.ModUserImg
	var FunctionSwitchs []mm.FunctionSwitch
	var UserInfoExts []mm.UserInfoExt
	var AddMsgs []mm.AddMsg

	if Response.CmdList != nil && len(Response.CmdList.List) > 0 {
		for _, v := range Response.CmdList.List {
			switch *v.CmdId {
			case int32(mm.SyncCmdID_CmdIdModUserInfo): // CmdId = 1
				var data mm.ModUserInfo
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				ModUserInfos = append(ModUserInfos, data)
			case int32(mm.SyncCmdID_CmdIdModContact): // CmdId = 2
				var data mm.ModContact
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				ModContacts = append(ModContacts, data)
			case int32(mm.SyncCmdID_CmdIdDelContact): // CmdId = 4
				var data mm.DelContact
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				DelContacts = append(DelContacts, data)
			case int32(mm.SyncCmdID_MM_SYNCCMD_MODUSERIMG): // CmdId = 35
				var data mm.ModUserImg
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				ModUserImgs = append(ModUserImgs, data)
			case int32(mm.SyncCmdID_CmdIdFunctionSwitch): // CmdId = 23
				var data mm.FunctionSwitch
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				FunctionSwitchs = append(FunctionSwitchs, data)
			case int32(mm.SyncCmdID_MM_SYNCCMD_USERINFOEXT): // CmdId = 44
				var data mm.UserInfoExt
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				UserInfoExts = append(UserInfoExts, data)
			case int32(mm.SyncCmdID_CmdIdAddMsg): // CmdId = 5
				var data mm.AddMsg
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				AddMsgs = append(AddMsgs, data)
			default:
				UnknownCmdId += UnknownCmdId + ";" + fmt.Sprintf("%v", *v.CmdId)
			}
		}
	}

	return wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data: Msg.SyncResponse{
			ModUserInfos:    ModUserInfos,
			ModContacts:     ModContacts,
			DelContacts:     DelContacts,
			ModUserImgs:     ModUserImgs,
			FunctionSwitchs: FunctionSwitchs,
			UserInfoExts:    UserInfoExts,
			AddMsgs:         AddMsgs,
			KeyBuf: mm.SKBuiltinBufferT{
				ILen:   Response.KeyBuf.ILen,
				Buffer: Response.KeyBuf.Buffer,
			},
			UnknownCmdId: UnknownCmdId,
			Remarks:      "出现未解析的CmdId类型数据,请联系客服人员处理。",
		},
	}
}
