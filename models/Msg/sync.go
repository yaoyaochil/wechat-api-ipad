package Msg

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type SyncParam struct {
	Wxid    string
	Scene   uint32
	Synckey string
}

type SyncResponse struct {
	ModUserInfos    []mm.ModUserInfo    //CmdId = 1
	ModContacts     []mm.ModContact     //CmdId = 2
	DelContacts     []mm.DelContact     //CmdId = 4
	ModUserImgs     []mm.ModUserImg     //CmdId = 35
	FunctionSwitchs []mm.FunctionSwitch //CmdId = 23
	UserInfoExts    []mm.UserInfoExt    //CmdId = 44
	AddMsgs         []mm.AddMsg         //CmdId = 5
	ContinueFlag    int32
	KeyBuf          mm.SKBuiltinBufferT
	Status          int32
	Continue        int32
	Time            int32
	UnknownCmdId    string
	Remarks         string
}

func Sync(Data SyncParam) wxCilent.ResponseResult {
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

	req := &mm.NewSyncRequest{
		Oplog: &mm.CmdList{
			Count: proto.Uint32(0),
			List:  nil,
		},
		Selector:      proto.Uint32(262151),
		KeyBuf:        &Synckey,
		Scene:         proto.Uint32(Data.Scene),
		DeviceType:    proto.String("iPhone"),
		SyncMsgDigest: proto.Uint32(3),
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
		Cgiurl:        "/cgi-bin/micromsg-bin/newsync",
		Proxy:         D.Proxy,
		Encryption:    0,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              138,
			Uin:              D.Uin,
			Cookie:           D.Cooike,
			Sessionkey:       D.Sessionkey,
			Loginecdhkey:     D.Loginecdhkey,
			Clientsessionkey: D.Clientsessionkey,
			Serversessionkey: D.Serversessionkey,
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
	NewSyncResponse := mm.NewSyncResponse{}
	err = proto.Unmarshal(protobufdata, &NewSyncResponse)
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

	if NewSyncResponse.CmdList != nil && len(NewSyncResponse.CmdList.List) > 0 {
		for _, v := range NewSyncResponse.CmdList.List {
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

		return wxCilent.ResponseResult{
			Code:    0,
			Success: true,
			Message: "成功",
			Data: SyncResponse{
				ModUserInfos:    ModUserInfos,
				ModContacts:     ModContacts,
				DelContacts:     DelContacts,
				ModUserImgs:     ModUserImgs,
				FunctionSwitchs: FunctionSwitchs,
				UserInfoExts:    UserInfoExts,
				AddMsgs:         AddMsgs,
				ContinueFlag:    *NewSyncResponse.ContinueFlag,
				KeyBuf: mm.SKBuiltinBufferT{
					ILen:   NewSyncResponse.KeyBuf.ILen,
					Buffer: NewSyncResponse.KeyBuf.Buffer,
				},
				Status:       *NewSyncResponse.Status,
				Continue:     *NewSyncResponse.Continue,
				Time:         *NewSyncResponse.Time,
				UnknownCmdId: UnknownCmdId,
				Remarks:      "",
			},
		}
	}

	return wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "当前未有新消息",
		Data:    NewSyncResponse,
	}
}
