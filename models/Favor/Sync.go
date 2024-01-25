package Favor

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

type SyncParam struct {
	Wxid   string
	Keybuf string
}

type SyncResponse struct {
	Ret    int32
	List   []mm.AddFavItem
	KeyBuf mm.SKBuiltinBufferT
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

	var KeyBuf mm.SKBuiltinBufferT

	if Data.Keybuf != "" {
		key, _ := base64.StdEncoding.DecodeString(Data.Keybuf)
		KeyBuf.Buffer = key
		KeyBuf.ILen = proto.Uint32(uint32(len(key)))
	}

	req := &mm.FavSyncRequest{
		Selector: proto.Uint32(1),
		KeyBuf:   &KeyBuf,
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
		Cgiurl:        "/cgi-bin/micromsg-bin/favsync",
		Proxy:         D.Proxy,
		Encryption:    5,
		TwelveEncData: wxCilent.PackSpecialCgiData{},
		PackData: wxCilent.PackData{
			Reqdata:          reqdata,
			Cgi:              400,
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
	Response := mm.FavSyncResponse{}
	err = proto.Unmarshal(protobufdata, &Response)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	var List []mm.AddFavItem

	for _, v := range Response.CmdList.List {
		if *v.CmdId == int32(mm.SyncCmdID_MM_FAV_SYNCCMD_ADDITEM) {
			var data mm.AddFavItem
			_ = proto.Unmarshal(v.CmdBuf.Buffer, &data)
			List = append(List, data)
		}
	}

	return wxCilent.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data: SyncResponse{
			Ret:    *Response.Ret,
			List:   List,
			KeyBuf: *Response.KeyBuf,
		},
	}

}
