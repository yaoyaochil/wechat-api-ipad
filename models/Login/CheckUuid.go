package Login

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

func CheckUuid(Uuid string) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Uuid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	timenow := uint32(time.Now().Unix())

	req := &mm.CheckLoginQRCodeRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Aeskey,
			Uin:           proto.Uint32(0),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		RandomEncryKey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(D.Aeskey))),
			Buffer: D.Aeskey,
		},
		Uuid:      &D.Uuid,
		TimeStamp: &timenow,
		Opcode:    proto.Uint32(0),
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

	//组包发包

	//开始请求发包
	protobufdata, cookie, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:         wxCilent.MMtls_ip,
		Host:       wxCilent.MMtls_host,
		Cgiurl:     "/cgi-bin/micromsg-bin/checkloginqrcode",
		Proxy:      D.Proxy,
		Encryption: 12,
		TwelveEncData: wxCilent.PackSpecialCgiData{
			Reqdata:                    reqdata,
			Cgi:                        502,
			Encrypttype:                12,
			Extenddata:                 []byte{},
			Uin:                        0,
			Cookies:                    []byte{},
			ClientVersion:              wxCilent.Wx_client_version,
			HybridEcdhPrivkey:          D.HybridEcdhPrivkey,
			HybridEcdhPubkey:           D.HybridEcdhPubkey,
			HybridEcdhInitServerPubKey: D.HybridEcdhInitServerPubKey,
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

	checkloginQRRes := mm.CheckLoginQRCodeResponse{}
	err = proto.Unmarshal(protobufdata, &checkloginQRRes)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	if checkloginQRRes.GetBaseResponse().GetRet() == 0 {
		if checkloginQRRes.GetNotifyPkg().GetNotifyData().GetBuffer() == nil {
			return wxCilent.ResponseResult{
				Code:    -8,
				Success: false,
				Message: "异常：扫码状态返回的交互key不存在",
				Data:    checkloginQRRes,
			}
		}

		notifydata := wxCilent.AesDecrypt(checkloginQRRes.GetNotifyPkg().GetNotifyData().GetBuffer(), D.NotifyKey)
		if notifydata != nil {
			notifydataRsp := mm.LoginQRCodeNotify{}
			err := proto.Unmarshal(notifydata, &notifydataRsp)
			if err != nil {
				return wxCilent.ResponseResult{
					Code:    -2,
					Success: false,
					Message: "解包异常",
					Data:    nil,
				}
			}

			//扫码确认登陆
			if notifydataRsp.GetStatus() == 2 {
				D.Wxid = notifydataRsp.GetUserName()
				D.Pwd = notifydataRsp.GetPwd()
				D.Cooike = cookie
				return CheckSecManualAuth(*D, wxCilent.MMtls_ip, wxCilent.MMtls_host)
			}

			return wxCilent.ResponseResult{
				Code:    0,
				Success: true,
				Message: "成功",
				Data:    notifydataRsp,
			}
		}
	}

	return wxCilent.ResponseResult{
		Code:    -0,
		Success: false,
		Message: "未知的错误",
		Data:    checkloginQRRes,
	}

}
