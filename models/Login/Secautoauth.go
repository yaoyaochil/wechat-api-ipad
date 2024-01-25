package Login

import (
	"crypto/md5"
	"fmt"
	"time"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/Cilent/device"
	"wechatwebapi/Cilent/mm"
	"wechatwebapi/comm"

	"github.com/golang/protobuf/proto"
)

func Secautoauth(Wxid string) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	if len(D.Autoauthkey) <= 0 {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "账号异常：Autoauthkey读取失败",
			Data:    nil,
		}
	}

	Autoauthkey := &mm.AutoAuthKey{}
	_ = proto.Unmarshal(D.Autoauthkey, Autoauthkey)

	Wx_login_prikey, Wx_login_pubkey := wxCilent.EcdhGen713Key()

	//基础设备信息
	Imei := device.Imei(D.Deviceid_str)
	SoftType := device.SoftType(D.Deviceid_str)
	ClientSeqId := wxCilent.GetClientSeqId(D.Deviceid_str)

	//24算法
	ccData := &mm.CryptoData{
		Version:     []byte("00000003"),
		Type:        proto.Uint32(1),
		EncryptData: wxCilent.GetNewSpamData(D.Deviceid_str, D.DeviceName),
		Timestamp:   proto.Uint32(uint32(time.Now().Unix())),
		Unknown5:    proto.Uint32(5),
		Unknown6:    proto.Uint32(0),
	}

	ccDataseq, _ := proto.Marshal(ccData)

	WCExtInfo := &mm.WCExtInfo{
		Wcstf: &mm.SKBuiltinBufferT{
			ILen:   nil,
			Buffer: nil,
		},
		Wcste: &mm.SKBuiltinBufferT{
			ILen:   nil,
			Buffer: nil,
		},
		CcData: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(ccDataseq))),
			Buffer: ccDataseq,
		},
	}

	WCExtInfoseq, _ := proto.Marshal(WCExtInfo)

	req := &mm.AutoAuthRequest{
		RsaReqData: &mm.AutoAuthRsaReqData{
			AesEncryptKey: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(Autoauthkey.EncryptKey.Buffer))),
				Buffer: Autoauthkey.EncryptKey.Buffer,
			},
			CliPubEcdhkey: &mm.ECDHKey{
				Nid: proto.Int32(713),
				Key: &mm.SKBuiltinBufferT{
					ILen:   proto.Uint32(uint32(len(Wx_login_pubkey))),
					Buffer: Wx_login_pubkey[:int32(len(Wx_login_pubkey))],
				},
			},
		},
		AesReqData: &mm.AutoAuthAesReqData{
			BaseRequest: &mm.BaseRequest{
				SessionKey:    D.Sessionkey,
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(wxCilent.Wx_client_version)),
				DeviceType:    wxCilent.DeviceType_byte,
				Scene:         proto.Uint32(0),
			},
			BaseReqInfo: &mm.BaseAuthReqInfo{},
			AutoAuthKey: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(D.Autoauthkey))),
				Buffer: D.Autoauthkey,
			},
			Imei:         &Imei,
			SoftType:     &SoftType,
			BuiltinIpseq: proto.Uint32(0),
			ClientSeqId:  &ClientSeqId,
			DeviceName:   proto.String(D.DeviceName),
			DeviceType:   proto.String("iPad"),
			Language:     proto.String("zh_CN"),
			TimeZone:     proto.String("8.0"),
			ExtSpamInfo: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(WCExtInfoseq))),
				Buffer: WCExtInfoseq,
			},
		},
	}

	reqdata, err := proto.Marshal(req)

	//开始发包请求
	protobufdata, cookie, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:         D.Mmtlsip,
		Host:       D.MmtlsHost,
		Cgiurl:     "/cgi-bin/micromsg-bin/secautoauth",
		Proxy:      D.Proxy,
		Encryption: 12,
		TwelveEncData: wxCilent.PackSpecialCgiData{
			Reqdata:                    reqdata,
			Cgi:                        763,
			Encrypttype:                12,
			Extenddata:                 []byte{},
			Uin:                        D.Uin,
			Cookies:                    D.Cooike,
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

	//解包
	UnifyAuthResponse := mm.UnifyAuthResponse{}
	err = proto.Unmarshal(protobufdata, &UnifyAuthResponse)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	loginRes := UnifyAuthResponse

	if loginRes.GetBaseResponse().GetRet() != 0 {
		return wxCilent.ResponseResult{
			Code:    int64(loginRes.GetBaseResponse().GetRet()),
			Success: false,
			Message: *loginRes.GetBaseResponse().GetErrMsg().String_,
			Data:    loginRes,
		}
	}

	Wx_loginecdhkey := wxCilent.DoECDH713(Wx_login_prikey, loginRes.GetAuthSectResp().GetSvrPubEcdhkey().GetKey().GetBuffer())
	Wx_loginecdhkeylen := int32(len(Wx_loginecdhkey))
	m := md5.New()
	m.Write(Wx_loginecdhkey[:Wx_loginecdhkeylen])
	D.Loginecdhkey = Wx_loginecdhkey
	ecdhdecrptkey := m.Sum(nil)
	D.Cooike = cookie
	D.Sessionkey = wxCilent.AesDecrypt(loginRes.GetAuthSectResp().GetSessionKey().GetBuffer(), ecdhdecrptkey)
	D.Autoauthkey = loginRes.GetAuthSectResp().GetAutoAuthKey().GetBuffer()
	D.Autoauthkeylen = int32(loginRes.GetAuthSectResp().GetAutoAuthKey().GetILen())
	D.Serversessionkey = loginRes.GetAuthSectResp().GetServerSessionKey().GetBuffer()
	D.Clientsessionkey = loginRes.GetAuthSectResp().GetClientSessionKey().GetBuffer()
	D.AuthTicket = loginRes.GetAuthSectResp().GetAuthTicket()

	err = comm.CreateLoginData(*D, D.Wxid, 0)

	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	return wxCilent.ResponseResult{
		Code:    1,
		Success: false,
		Message: "登陆成功",
		Data:    loginRes,
	}
}
