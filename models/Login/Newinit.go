package Login

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type newinitResponse struct {
	BaseResponse    mm.BaseResponse
	CurrentSynckey  mm.SKBuiltinBufferT
	MaxSynckey      mm.SKBuiltinBufferT
	ContinueFlag    uint32
	SelectBitmap    uint32
	ModUserInfos    []mm.ModUserInfo    //CmdId = 1
	ModUserImgs     []mm.ModUserImg     //CmdId = 35
	UserInfoExts    []mm.UserInfoExt    //CmdId = 44
	FunctionSwitchs []mm.FunctionSwitch //CmdId = 23
	UnknownCmdId    string
	Remarks         string
}

func Newinit(Wxid, cursync, maxsync string) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	Cursync := []byte(cursync)
	Maxsync := []byte(maxsync)

	if cursync == "" {
		Cursync = make([]byte, 0)
	}

	if maxsync == "" {
		Maxsync = make([]byte, 0)
	}

	req := &mm.NewInitRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		UserName: &D.Wxid,
		CurrentSynckey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(Cursync))),
			Buffer: Cursync,
		},
		MaxSynckey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(Maxsync))),
			Buffer: Maxsync,
		},
		Language: proto.String("zh"),
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
		Cgiurl:        "/cgi-bin/micromsg-bin/newinit",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              139,
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
	NewInitResponse := mm.NewInitResponse{}
	err = proto.Unmarshal(protobufdata, &NewInitResponse)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	var ModUserInfos []mm.ModUserInfo
	var ModUserImgs []mm.ModUserImg
	var FunctionSwitchs []mm.FunctionSwitch
	var UserInfoExts []mm.UserInfoExt

	UnknownCmdId := ""

	if NewInitResponse.CmdList != nil && len(NewInitResponse.CmdList) > 0 {
		for _, v := range NewInitResponse.CmdList {
			switch *v.CmdId {
			case int32(mm.SyncCmdID_CmdIdModUserInfo): // CmdId = 1
				var data mm.ModUserInfo
				_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
				ModUserInfos = append(ModUserInfos, data)
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
			default:
				UnknownCmdId += fmt.Sprintf("%v", *v.CmdId) + ";"
			}
		}

		return wxCilent.ResponseResult{
			Code:    0,
			Success: true,
			Message: "成功",
			Data: newinitResponse{
				BaseResponse:    *NewInitResponse.BaseResponse,
				CurrentSynckey:  *NewInitResponse.CurrentSynckey,
				MaxSynckey:      *NewInitResponse.MaxSynckey,
				ContinueFlag:    *NewInitResponse.ContinueFlag,
				SelectBitmap:    *NewInitResponse.SelectBitmap,
				ModUserInfos:    ModUserInfos,
				ModUserImgs:     ModUserImgs,
				UserInfoExts:    UserInfoExts,
				FunctionSwitchs: FunctionSwitchs,
				UnknownCmdId:    UnknownCmdId,
				Remarks:         "出现未解析的CmdId类型数据,请联系客服人员处理。",
			},
		}
	}

	return wxCilent.ResponseResult{
		Code:    -8,
		Success: false,
		Message: "失败：未知原因",
		Data:    NewInitResponse,
	}

}
