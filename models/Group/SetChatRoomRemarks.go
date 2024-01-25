package Group

import (
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/bts"
	"wechatwebapi/comm"
	"wechatwebapi/models/Tools"

	"github.com/golang/protobuf/proto"
)

func SetChatRoomRemarks(Data OperateChatRoomInfoParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	GetContact := Tools.GetContact(Tools.GetContactParam{
		Wxid:         Data.Wxid,
		UserNameList: Data.QID,
	})

	if GetContact.Data == nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", GetContact.Message),
			Data:    nil,
		}
	}

	Contact := bts.GetContactResponse(GetContact.Data)

	ModContact := &mm.ModContact{
		UserName: &mm.SKBuiltinStringT{
			String_: proto.String(Data.QID),
		},
		NickName:  &mm.SKBuiltinStringT{},
		PyInitial: &mm.SKBuiltinStringT{},
		QuanPin:   &mm.SKBuiltinStringT{},
		Sex:       proto.Int32(0),
		ImgBuf:    &mm.SKBuiltinBufferT{},
		BitMask:   Contact.ContactList[0].BitMask,
		BitVal:    proto.Uint32(32770),
		ImgFlag:   proto.Uint32(0),
		Remark: &mm.SKBuiltinStringT{
			String_: proto.String(Data.Content),
		},
		RemarkPyinitial: &mm.SKBuiltinStringT{},
		RemarkQuanPin:   &mm.SKBuiltinStringT{},
		ContactType:     proto.Uint32(0),
		ChatRoomNotify:  proto.Uint32(1),
		AddContactScene: proto.Uint32(0),
		Extflag:         proto.Int32(0),
	}

	buffer, err := proto.Marshal(ModContact)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	cmdItem := mm.CmdItem{
		CmdId: proto.Int32(2),
		CmdBuf: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(buffer))),
			Buffer: buffer,
		},
	}

	var cmdItems []*mm.CmdItem
	cmdItems = append(cmdItems, &cmdItem)

	req := &mm.OpLogRequest{
		Cmd: &mm.CmdList{
			Count: proto.Uint32(uint32(len(cmdItems))),
			List:  cmdItems,
		},
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
	GetContactResponse := mm.OplogResponse{}
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
