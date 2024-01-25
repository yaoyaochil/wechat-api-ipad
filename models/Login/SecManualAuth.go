package Login

import (
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/device"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"
)

func SecManualAuth(Data comm.LoginData, mmtlsip, mmtlshost string) (mm.UnifyAuthResponse, []byte, []byte) {
	Wx_login_prikey, Wx_login_pubkey := wxCilent.EcdhGen713Key()
	var Wx_login_pubkeylen, Wx_login_prikeylen int32
	if len(Wx_login_prikey) > 0 && len(Wx_login_pubkey) > 0 {
		Wx_login_pubkeylen = int32(len(Wx_login_pubkey))
		Wx_login_prikeylen = int32(len(Wx_login_prikey))
	}

	log.Info(Wx_login_prikeylen)

	aeskey := []byte(wxCilent.RandSeq(16)) //获取随机密钥
	accountRequest := &mm.ManualAuthRsaReqData{
		RandomEncryKey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(aeskey))),
			Buffer: aeskey,
		},
		CliPubEcdhkey: &mm.ECDHKey{
			Nid: proto.Int32(713),
			Key: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(Wx_login_pubkeylen)),
				Buffer: Wx_login_pubkey[:Wx_login_pubkeylen],
			},
		},
		UserName: &Data.Wxid,
		Pwd:      &Data.Pwd,
		Pwd2:     &Data.Pwd,
	}
	ccData := &mm.CryptoData{
		Version:     []byte("00000003"),
		Type:        proto.Uint32(1),
		EncryptData: wxCilent.GetNewSpamData(Data.Deviceid_str, "ipad7,5"),
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
	ClientSeqId := wxCilent.GetClientSeqId(Data.Deviceid_str)
	Imei := device.Imei(Data.Deviceid_str)
	SoftType := device.SoftType(Data.Deviceid_str)
	uuid1, _ := device.Uuid(Data.Deviceid_str)

	deviceRequest := &mm.ManualAuthAesReqData{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    aeskey,
			Uin:           proto.Uint32(0),
			DeviceId:      Data.Deviceid_byte,
			ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
			DeviceType:    wxCilent.DeviceType_byte,
			Scene:         proto.Uint32(1),
		},
		BaseReqInfo:  &mm.BaseAuthReqInfo{},
		Imei:         &Imei,
		SoftType:     &SoftType,
		BuiltinIpseq: proto.Uint32(0),
		ClientSeqId:  &ClientSeqId,
		DeviceName:   proto.String("iPad7,5"),
		DeviceType:   proto.String("iPad"),
		Language:     proto.String("zh_CN"),
		TimeZone:     proto.String("8.0"),
		Channel:      proto.Int(0),
		TimeStamp:    proto.Uint32(uint32(time.Now().Unix())),
		DeviceBrand:  proto.String("Apple"),
		Ostype:       &wxCilent.DeviceType_str,
		RealCountry:  proto.String("CN"),
		BundleId:     proto.String("com.tencent.xin"),
		AdSource:     &uuid1,
		IphoneVer:    proto.String("iPad11,3"),
		InputType:    proto.Uint32(2),
		ExtSpamInfo: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(WCExtInfoseq))),
			Buffer: WCExtInfoseq,
		},
	}

	requset := &mm.SecManualLoginRequest{
		RsaReqData: accountRequest,
		AesReqData: deviceRequest,
	}
	reqdata, err := proto.Marshal(requset)

	//开始发包请求
	protobufdata, cookie, _, err := comm.SendRequest(comm.SendPostData{
		Ip:         mmtlsip,
		Host:       mmtlshost,
		Cgiurl:     "/cgi-bin/micromsg-bin/secmanualauth",
		Proxy:      Data.Proxy,
		Encryption: 12,
		TwelveEncData: wxCilent.PackSpecialCgiData{
			Reqdata:                    reqdata,
			Cgi:                        252,
			Encrypttype:                12,
			Extenddata:                 []byte{},
			Uin:                        0,
			Cookies:                    []byte{},
			ClientVersion:              wxCilent.Wx_client_version,
			HybridEcdhPrivkey:          Data.HybridEcdhPrivkey,
			HybridEcdhPubkey:           Data.HybridEcdhPubkey,
			HybridEcdhInitServerPubKey: Data.HybridEcdhInitServerPubKey,
		},
	}, Data.MmtlsKey)

	loginRes := mm.UnifyAuthResponse{}
	err = proto.Unmarshal(protobufdata, &loginRes)
	if err == nil {
		return loginRes, Wx_login_prikey, cookie
	}

	return loginRes, []byte{}, []byte{}
}
