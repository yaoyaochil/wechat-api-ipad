package Cilent

import "hash"

// 0x17000841 708
// 0x17000C2B 712

var Wx_client_version = 0x17000C2B
var MMtls_ip = "szshort.weixin.qq.com"
var MMtls_host = "extshort.weixin.qq.com"
var DeviceType_byte = []byte("iPad iOS13.3.1")
var DeviceType_str = "iPad iOS13.3.1"

var HybridDecryptHash hash.Hash
var HybridServerpubhashFinal hash.Hash

type PackSpecialCgiData struct {
	Reqdata                    []byte
	Cgi                        int
	Encrypttype                int
	Extenddata                 []byte
	Uin                        uint32
	Cookies                    []byte
	ClientVersion              int
	HybridEcdhPrivkey          []byte
	HybridEcdhPubkey           []byte
	HybridEcdhInitServerPubKey []byte
}

type PackData struct {
	Reqdata          []byte
	Cgi              int
	Uin              uint32
	Cookie           []byte
	ClientVersion    int
	Sessionkey       []byte
	EncryptType      uint8
	Loginecdhkey     []byte
	Clientsessionkey []byte
	Serversessionkey []byte
	UseCompress      bool
}

type ResponseResult struct {
	Code    int64
	Success bool
	Message string
	Data    interface{}
}
