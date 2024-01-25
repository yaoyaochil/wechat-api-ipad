package Login

import (
	"container/list"
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"strings"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/comm"
)

func CheckSecManualAuth(Data comm.LoginData, mmtlsip, mmtlshost string) wxCilent.ResponseResult {
	//开始登陆
	loginRes, Wx_login_prikey, Cookie := SecManualAuth(Data, mmtlsip, mmtlshost)

	//登陆成功
	if loginRes.GetBaseResponse().GetRet() == 0 && loginRes.GetUnifyAuthSectFlag() > 0 {
		Wx_loginecdhkey := wxCilent.DoECDH713(Wx_login_prikey, loginRes.GetAuthSectResp().GetSvrPubEcdhkey().GetKey().GetBuffer())
		Wx_loginecdhkeylen := int32(len(Wx_loginecdhkey))
		m := md5.New()
		m.Write(Wx_loginecdhkey[:Wx_loginecdhkeylen])

		Data.Loginecdhkey = Wx_loginecdhkey

		ecdhdecrptkey := m.Sum(nil)
		Data.Uin = loginRes.GetAuthSectResp().GetUin()
		Data.Wxid = loginRes.GetAcctSectResp().GetUserName()
		Data.Alais = loginRes.GetAcctSectResp().GetAlias()
		Data.Mobile = loginRes.GetAcctSectResp().GetBindMobile()
		Data.NickName = loginRes.GetAcctSectResp().GetNickName()
		Data.Cooike = Cookie
		Data.Sessionkey = wxCilent.AesDecrypt(loginRes.GetAuthSectResp().GetSessionKey().GetBuffer(), ecdhdecrptkey)
		Data.Sessionkey_2 = loginRes.GetAuthSectResp().GetSessionKey().GetBuffer()
		Data.Autoauthkey = loginRes.GetAuthSectResp().GetAutoAuthKey().GetBuffer()
		Data.Autoauthkeylen = int32(loginRes.GetAuthSectResp().GetAutoAuthKey().GetILen())
		Data.Serversessionkey = loginRes.GetAuthSectResp().GetServerSessionKey().GetBuffer()
		Data.Clientsessionkey = loginRes.GetAuthSectResp().GetClientSessionKey().GetBuffer()
		Data.AuthTicket = loginRes.GetAuthSectResp().GetAuthTicket()
		Data.Mmtlsip = mmtlsip
		Data.MmtlsHost = mmtlshost
		Data.ClientVersion = wxCilent.Wx_client_version
		Data.DeviceType = wxCilent.DeviceType_str

		err := comm.CreateLoginData(Data, Data.Wxid, 0)

		if err != nil {
			return wxCilent.ResponseResult{
				Code:    -8,
				Success: false,
				Message: fmt.Sprintf("系统异常：%v", err.Error()),
				Data:    nil,
			}
		}

		type Suc struct {
			Wxid     string
			Alais    string
			NickName string
			Mobile   string
		}

		return wxCilent.ResponseResult{
			Code:    0,
			Success: true,
			Message: "登陆成功",
			Data:    loginRes,
		}
	}

	//30系列转向
	if loginRes.GetBaseResponse().GetRet() == -301 {
		var Wx_newLongIPlist, Wx_newshortIplist, Wx_newshortextipList list.List
		var Wx_newLong_Host, Wx_newshort_Host, Wx_newshortext_Host list.List

		dns_info := loginRes.GetNetworkSectResp().GetNewHostList().GetList()
		for _, v := range dns_info {
			if v.GetHost() == "long.weixin.qq.com" {
				ip_info := loginRes.GetNetworkSectResp().GetBuiltinIplist().GetLongConnectIplist()
				for _, ip := range ip_info {
					host := ip.GetHost()
					host = strings.Replace(host, string(byte(0x00)), "", -1)
					if host == v.GetRedirect() {
						ipaddr := ip.GetIp()
						ipaddr = strings.Replace(ipaddr, string(byte(0x00)), "", -1)
						Wx_newLongIPlist.PushBack(ipaddr)
						Wx_newLong_Host.PushBack(host)
					}
				}
			} else if v.GetHost() == "short.weixin.qq.com" {
				ip_info := loginRes.GetNetworkSectResp().GetBuiltinIplist().GetShortConnectIplist()
				for _, ip := range ip_info {
					host := ip.GetHost()
					host = strings.Replace(host, string(byte(0x00)), "", -1)
					if host == v.GetRedirect() {
						ipaddr := ip.GetIp()
						ipaddr = strings.Replace(ipaddr, string(byte(0x00)), "", -1)
						Wx_newshortIplist.PushBack(ipaddr)
						Wx_newshort_Host.PushBack(host)
					}
				}
			} else if v.GetHost() == "extshort.weixin.qq.com" {
				ip_info := loginRes.GetNetworkSectResp().GetBuiltinIplist().GetShortConnectIplist()
				for _, ip := range ip_info {
					host := ip.GetHost()
					host = strings.Replace(host, string(byte(0x00)), "", -1)
					if host == v.GetRedirect() {
						ipaddr := ip.GetIp()
						ipaddr = strings.Replace(ipaddr, string(byte(0x00)), "", -1)
						Wx_newshortextipList.PushBack(ipaddr)
						Wx_newshortext_Host.PushBack(host)
					}
				}
			}
		}
		return CheckSecManualAuth(Data, Wx_newshortIplist.Front().Value.(string), Wx_newshort_Host.Front().Value.(string))
	}

	//否则就是包有问题
	return wxCilent.ResponseResult{
		Code:    int64(loginRes.GetBaseResponse().GetRet()),
		Success: false,
		Message: filterRetMessage(*loginRes.GetBaseResponse().GetErrMsg().String_),
		Data:    loginRes,
	}
}

type RetMsg struct {
	content string `xml:"Content"`
	url     string `xml:"Url"`
}

type Root struct {
	msg RetMsg `xml:"e"`
}

func filterRetMessage(Message string) string {
	if strings.Contains(Message, "使用存在异常") {
		return "当前帐号的使用存在异常，为保护帐号安全，系统将其自动置为保护状态，被限制登录，如需继续使用，请轻触“了解详情”申请解除限制。"
	} else if strings.Contains(Message, "环境存在异常") {
		return "系统检测到环境存在异常，为了你的帐号安全， 请轻触“确定”进行安全验证"
	} else if strings.Contains(Message, "恶意营销") {
		return "该微信帐号因存在骚扰/恶意营销/欺诈等违规行为被限制登录，如需继续使用，请轻触“了解详情”申请解除限制。"
	} else if strings.Contains(Message, "外挂") {
		return "该微信帐号因使用了微信外挂、非官方客户端或模拟器，被永久限制登录，请尽快卸载对应的非法软件。若帐号内有资金，可轻触“了解详情”按相关指引进行操作。"
	} else if strings.Contains(Message, "自助冻结") {
		return "微信帐号已通过手机自助冻结，如确认帐号当前处于安全状态，可以点击确定按钮解冻"
	} else if strings.Contains(Message, "密码错误") {
		return "帐号或密码错误，请重新填写。"
	} else {
		return Message
	}
}

func filterRetMessage2(Message string) string {
	r := Root{}
	err := xml.Unmarshal([]byte(Message), &r)
	if err != nil {
		fmt.Printf("error: %v", err)
		return Message
	} else {
		return r.msg.content + " " + r.msg.url
	}
}
