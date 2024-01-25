package Login

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
	"wechatwebapi/models"
)

type GetQRReq struct {
	Proxy      models.ProxyInfo
	DeviceID   string
	DeviceName string
}

type GetQRRes struct {
	QrBase64    string
	Uuid        string
	QrUrl       string
	ExpiredTime string
}

func GetQRCODE(DeviceID, DeviceName string, Proxy models.ProxyInfo) wxCilent.ResponseResult {
	//初始化Mmtls
	_, MmtlsClient, err := comm.MmtlsInitialize(Proxy)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	aeskey := []byte(wxCilent.RandSeq(16)) //获取随机密钥
	deviceid := wxCilent.CreateDeviceId(DeviceID)
	devicelIdByte, _ := hex.DecodeString(deviceid)

	HybridEcdhInitServerPubKey, HybridEcdhPrivkey, HybridEcdhPubkey := wxCilent.HybridEcdhInit()

	req := &mm.GetLoginQRCodeRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    aeskey,
			Uin:           proto.Uint32(0),
			DeviceId:      devicelIdByte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(0),
		},
		RandomEncryKey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(aeskey))),
			Buffer: aeskey,
		},
		Opcode:           proto.Uint32(0),
		MsgContextPubKey: nil,
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

	//开始请求发包
	protobufdata, cookie, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:         wxCilent.MMtls_ip,
		Host:       wxCilent.MMtls_host,
		Cgiurl:     "/cgi-bin/micromsg-bin/getloginqrcode",
		Proxy:      Proxy,
		Encryption: 12,
		TwelveEncData: wxCilent.PackSpecialCgiData{
			Reqdata:                    reqdata,
			Cgi:                        501,
			Encrypttype:                12,
			Extenddata:                 []byte{},
			Uin:                        0,
			Cookies:                    []byte{},
			ClientVersion:              wxCilent.Wx_client_version,
			HybridEcdhPrivkey:          HybridEcdhPrivkey,
			HybridEcdhPubkey:           HybridEcdhPubkey,
			HybridEcdhInitServerPubKey: HybridEcdhInitServerPubKey,
		},
	}, MmtlsClient)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	getloginQRRes := mm.GetLoginQRCodeResponse{}

	err = proto.Unmarshal(protobufdata, &getloginQRRes)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	if getloginQRRes.GetBaseResponse().GetRet() == 0 {
		//Wx_qrcode_uuid = getloginQRRes.GetUuid()
		//Wx_NotifyKey = getloginQRRes.GetNotifyKey().GetBuffer()
		//Wx_qrimagecode = getloginQRRes.GetQrcode().GetBuffer()

		//保存redis
		err := comm.CreateLoginData(comm.LoginData{
			Uuid:                       getloginQRRes.GetUuid(),
			Aeskey:                     aeskey,
			NotifyKey:                  getloginQRRes.GetNotifyKey().GetBuffer(),
			Deviceid_str:               deviceid,
			Deviceid_byte:              devicelIdByte,
			DeviceName:                 DeviceName,
			HybridEcdhPrivkey:          HybridEcdhPrivkey,
			HybridEcdhPubkey:           HybridEcdhPubkey,
			HybridEcdhInitServerPubKey: HybridEcdhInitServerPubKey,
			Cooike:                     cookie,
			Proxy:                      Proxy,
			MmtlsKey:                   MmtlsClient,
		}, "", 300)

		if err == nil {
			return wxCilent.ResponseResult{
				Code:    1,
				Success: true,
				Message: "成功",
				Data: GetQRRes{
					"",
					getloginQRRes.GetUuid(),
					"http://qr.topscan.com/api.php?text=http://weixin.qq.com/x/" + getloginQRRes.GetUuid(),
					time.Unix(int64(getloginQRRes.GetExpiredTime()), 0).Format("2006-01-02 15:04:05"),
				},
			}
		}
	}

	return wxCilent.ResponseResult{
		Code:    -0,
		Success: false,
		Message: "未知的错误",
		Data:    getloginQRRes,
	}
}
