package Mmtls

import (
	"crypto/sha256"
	"net/http"
)

type MmtlsClient struct {
	Shakehandpubkey    []byte
	Shakehandpubkeylen int32
	Shakehandprikey    []byte
	Shakehandprikeylen int32

	Shakehandpubkey_2   []byte
	Shakehandpubkeylen2 int32
	Shakehandprikey_2   []byte
	Shakehandprikeylen2 int32

	Mserverpubhashs     []byte
	ServerSeq           int
	ClientSeq           int
	ShakehandECDHkey    []byte
	ShakehandECDHkeyLen int32

	Encrptmmtlskey  []byte
	Decryptmmtlskey []byte
	EncrptmmtlsIv   []byte
	DecryptmmtlsIv  []byte

	CurDecryptSeqIv []byte
	CurEncryptSeqIv []byte

	Decrypt_part2_hash256            []byte
	Decrypt_part3_hash256            []byte
	ShakehandECDHkeyhash             []byte
	Hkdfexpand_pskaccess_key         []byte
	Hkdfexpand_pskrefresh_key        []byte
	HkdfExpand_info_serverfinish_key []byte
	Hkdfexpand_clientfinish_key      []byte
	Hkdfexpand_secret_key            []byte

	Hkdfexpand_application_key []byte
	Encrptmmtlsapplicationkey  []byte
	Decryptmmtlsapplicationkey []byte
	EncrptmmtlsapplicationIv   []byte
	DecryptmmtlsapplicationIv  []byte

	Earlydatapart       []byte
	Newsendbufferhashs  []byte
	Encrptshortmmtlskey []byte
	Encrptshortmmtlsiv  []byte
	Decrptshortmmtlskey []byte
	Decrptshortmmtlsiv  []byte

	//http才需要
	Pskkey    string
	Pskiv     string
	MmtlsMode uint
}

//短连接mmtls模块
type HttpClientModel struct {
	mmtlsClient *MmtlsClient
	httpClient  *http.Client
	curShortip  string
	mmtlsIsInit bool
}

type MmtlsPacketHeader struct {
	headerbyte      byte
	headerversion   uint16
	headerpacketLen uint16
}

func Getsha256(Data []byte) []byte {
	D := sha256.New()
	D.Write(Data)
	return D.Sum(nil)
}