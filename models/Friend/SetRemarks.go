package Friend

import (
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/bts"
	"wechatwebapi/comm"

	"github.com/golang/protobuf/proto"
)

type SetRemarksParam struct {
	Wxid    string
	ToWxids []string
	Remarks string
}

func SetRemarks(Data SetRemarksParam) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	//先获取用户的基本信息
	getContact := GetContact(GetContactDetailparameter{
		Wxid:     Data.Wxid,
		Towxids:  Data.ToWxids,
		ChatRoom: "",
	})

	if getContact.Code != 0 {
		return getContact
	}

	Contact := bts.GetContactResponse(getContact.Data)

	if len(Contact.ContactList) > 0 {
		modContact := Contact.ContactList[0]
		modContact.Remark = &mm.SKBuiltinStringT{
			String_: proto.String(Data.Remarks),
		}

		ContactList := &mm.ModContact{
			UserName:        modContact.UserName,
			NickName:        modContact.NickName,
			PyInitial:       modContact.PyInitial,
			QuanPin:         modContact.QuanPin,
			Sex:             modContact.Sex,
			ImgBuf:          modContact.ImgBuf,
			BitMask:         modContact.BitMask,
			BitVal:          modContact.BitVal,
			ImgFlag:         modContact.ImgFlag,
			Remark:          modContact.Remark,
			RemarkPyinitial: modContact.RemarkPyinitial,
			RemarkQuanPin:   modContact.RemarkQuanPin,
			ContactType:     modContact.ContactType,
			ChatRoomNotify:  proto.Uint32(1),
			AddContactScene: modContact.AddContactScene,
			Extflag:         modContact.Extflag,
		}

		var cmdItems []*mm.CmdItem
		buffer, err := proto.Marshal(ContactList)
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

	return wxCilent.ResponseResult{}
}
