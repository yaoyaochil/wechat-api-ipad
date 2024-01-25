package Cilent

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"strings"
	"wechatwebapi/Cilent/mm"
)

func HybridEcdhInit() (HybridEcdhInitServerPubKey, HybridEcdhPrivkey, HybridEcdhPubkey []byte) {
	wxHybridEcdhPubKey := "04 7e be 76 04 ac f0 72 b0 ab 01 77 ea 55 1a 7b 72 58 8f 9b 5d 38 01 df d7 bb 1b ca 8e 33 d1 c3" +
		"b8 fa 6e 4e 40 26 eb 38 d5 bb 36 50 88 a3 d3 16" +
		"7c 83 bd d0 bb b4 62 55 f8 8a 16 ed e6 f7 ab 43" +
		"b5"
	strpubkey := strings.Replace(wxHybridEcdhPubKey, " ", "", -1)
	HybridEcdhInitServerPubKey, _ = hex.DecodeString(strpubkey)
	HybridEcdhPrivkey, HybridEcdhPubkey = GenECDH415Key()
	return
}

type HybridEcdhEncryptData struct {
	Src []byte
	Externkey []byte
	HybridEcdhPrivkey []byte
	HybridEcdhPubkey  []byte
	HybridEcdhInitServerPubKey []byte
}

func GenECDH415Key() (privKey []byte, pubKey []byte) {
	privKey = nil
	pubKey = nil
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := &priv.PublicKey
	pubKey = elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	privKey = priv.D.Bytes()
	return
}

func DoECDH415(privD, pubData []byte) []byte {
	X, Y := elliptic.Unmarshal(elliptic.P256(), pubData)
	if X == nil || Y == nil {
		log.Println("if X == nil || Y == nil ")
		return nil
	}
	x, _ := elliptic.P256().ScalarMult(X, Y, privD)
	return x.Bytes()
}

func HybridHkdfExpand(prikey []byte, salt []byte, info []byte, outLen int) []byte {
	h := hmac.New(sha256.New, prikey)
	h.Write(salt)
	T := h.Sum(nil)
	return HkdfExpand(sha256.New, T, info, outLen)
}


func HybridEcdhEncrypt(Data HybridEcdhEncryptData) (ret_bytes []byte) {
	ecdhkey := DoECDH415(Data.HybridEcdhPrivkey, Data.HybridEcdhInitServerPubKey)
	m := sha256.New()
	m.Write(ecdhkey[:])
	ecdhkey = m.Sum(nil)
	mClientpubhash := sha256.New()
	mClientpubhash.Write([]byte("1"))
	mClientpubhash.Write([]byte("415"))
	mClientpubhash.Write(Data.HybridEcdhPubkey)
	mClientpubhash_digest := mClientpubhash.Sum(nil)

	mRandomEncryptKey := []byte(RandSeq(32)) //获取随机密钥

	mNonce := []byte(RandSeq(12)) //获取随机密钥
	mEncryptdata := AesGcmEncryptWithCompress(ecdhkey[:24], mRandomEncryptKey, mNonce, mClientpubhash_digest)
	var mExternEncryptdata []byte
	if len(Data.Externkey) == 0x20 {
		mExternEncryptdata = AesGcmEncryptWithCompress(Data.Externkey[:24], mRandomEncryptKey, mNonce, mClientpubhash_digest)
	}
	hkdfexpand_security_key := HybridHkdfExpand([]byte("security hdkf expand"), mRandomEncryptKey, mClientpubhash_digest, 56)

	mClientpubhashFinal := sha256.New()
	mClientpubhashFinal.Write([]byte("1"))
	mClientpubhashFinal.Write([]byte("415"))
	mClientpubhashFinal.Write(Data.HybridEcdhPubkey)
	mClientpubhashFinal.Write(mEncryptdata)
	mClientpubhashFinal.Write(mExternEncryptdata)
	mClientpubhashFinal_digest := mClientpubhashFinal.Sum(nil)

	mEncryptdataFinal := AesGcmEncryptWithCompress(hkdfexpand_security_key[:24], Data.Src, mNonce, mClientpubhashFinal_digest)

	HybridDecryptHash = sha256.New()
	HybridDecryptHash.Write(mEncryptdataFinal)

	HybridServerpubhashFinal = sha256.New()
	HybridServerpubhashFinal.Write(hkdfexpand_security_key[24:56])
	HybridServerpubhashFinal.Write(Data.Src)

	HybridEcdhRequest := &mm.HybridEcdhRequest{
		Type: proto.Int32(1),
		SecECDHKey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(415),
			Buffer: Data.HybridEcdhPubkey,
		},
		Randomkeydata:       mEncryptdata,
		Randomkeyextenddata: mExternEncryptdata,
		Encyptdata:          mEncryptdataFinal,
	}
	reqdata, _ := proto.Marshal(HybridEcdhRequest)
	return reqdata
}

func  HybridEcdhDecrypt(src_bytes []byte, HybridEcdhPrivkey []byte) (ret_bytes []byte) {
	HybridEcdhResponse := &mm.HybridEcdhResponse{}
	err := proto.Unmarshal(src_bytes, HybridEcdhResponse)
	if err == nil {
		decrptecdhkey := DoECDH415(HybridEcdhPrivkey, HybridEcdhResponse.SecECDHKey.Buffer)
		m := sha256.New()
		m.Write(decrptecdhkey[:])
		decrptecdhkey = m.Sum(nil)
		HybridServerpubhashFinal.Write([]byte("415"))
		HybridServerpubhashFinal.Write(HybridEcdhResponse.SecECDHKey.Buffer)
		HybridServerpubhashFinal.Write([]byte("1"))
		mServerpubhashFinal_digest := HybridServerpubhashFinal.Sum(nil)

		outdata := AesGcmDecryptWithUncompress(decrptecdhkey[:24], HybridEcdhResponse.Decryptdata, mServerpubhashFinal_digest)
		return outdata
	}
	return nil
}

func UnpackBusinessHybridEcdhPacket(data []byte, uin uint32, cookies *[]byte, HybridEcdhPrivkey []byte) []byte {
	var body []byte
	if len(data) < 0x20 {
		return nil
	} else {
		var nCur int64
		var bfbit byte
		srcreader := bytes.NewReader(data)
		binary.Read(srcreader, binary.BigEndian, &bfbit)
		if bfbit == byte(0xbf) {
			nCur += 1
		}
		nLenHeader := data[nCur] >> 2
		nCur += 1
		nLenCookie := data[nCur] & 0xf
		nCur += 1
		nCur += 4
		srcreader.Seek(nCur, io.SeekStart)
		binary.Read(srcreader, binary.BigEndian, &uin)
		nCur += 4
		cookie_temp := data[nCur : nCur+int64(nLenCookie)]
		*cookies = cookie_temp
		nCur += int64(nLenCookie)
		cgidata := data[nCur:]
		_, nSize := proto.DecodeVarint(cgidata)
		nCur += int64(nSize)
		LenProtobufData := data[nCur:]
		_, nLenProtobuf := proto.DecodeVarint(LenProtobufData)
		nCur += int64(nLenProtobuf)
		body = data[nLenHeader:]
	}
	protobufdata := HybridEcdhDecrypt(body,HybridEcdhPrivkey)
	return protobufdata
}


func DoECDH713(privD, pubData []byte) []byte {
	X, Y := elliptic.Unmarshal(elliptic.P224(), pubData)
	if X == nil || Y == nil {
		return *new([]byte)
	}
	x, _ := elliptic.P224().ScalarMult(X, Y, privD)
	return x.Bytes()

}

func EcdhGen713Key() (privKey []byte, pubKey []byte) {
	privKey = nil
	pubKey = nil
	priv, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	pub := &priv.PublicKey
	pubKey = elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	privKey = priv.D.Bytes()
	return
}
