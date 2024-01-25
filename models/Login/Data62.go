package Login

import (
	"encoding/hex"
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/device"
	"wechatwebapi/comm"
	"wechatwebapi/models"
)

type ManualAuthReq struct {
	UserName   string
	Password   string
	DeviceData string
	AndroidDevice string
	AndroidOstype string
	AndroidBuild string
	AndroidBaseBand string
	AndroidID string
	Mac string
	Iccid string
	Imsi string
	PhoneSerial string
	Proxy      models.ProxyInfo
}

func Data62Login(D ManualAuthReq) wxCilent.ResponseResult {
	if D.UserName == "" || D.Password == "" {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "请输入账号或密码",
			Data:    nil,
		}
	}

	//初始化Mmtls
	_, MmtlsClient, err := comm.MmtlsInitialize(D.Proxy)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	devicelId := wxCilent.Get62Key(D.DeviceData)
	println(devicelId)
	devicelId = "49" + devicelId[2:]
	println(devicelId)
	devicelIdByte, _ := hex.DecodeString(devicelId)

	HybridEcdhInitServerPubKey, HybridEcdhPrivkey, HybridEcdhPubkey := wxCilent.HybridEcdhInit()

	LoginData := comm.LoginData{
		Wxid:                       D.UserName,
		Pwd:                        device.MD5ToLower(D.Password),
		Deviceid_str:               devicelId,
		Deviceid_byte:              devicelIdByte,
		DeviceName:                 "iPad7,5",
		Mmtlsip:                    wxCilent.MMtls_ip,
		MmtlsHost:                  wxCilent.MMtls_host,
		HybridEcdhPrivkey:          HybridEcdhPrivkey,
		HybridEcdhPubkey:           HybridEcdhPubkey,
		HybridEcdhInitServerPubKey: HybridEcdhInitServerPubKey,
		Proxy:                      D.Proxy,
		MmtlsKey:                   MmtlsClient,
	}

	return CheckSecManualAuth(LoginData, wxCilent.MMtls_ip, wxCilent.MMtls_host)
}
