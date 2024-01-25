package Mmtls

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"strconv"
	"time"
	"wechatwebapi/Cilent"
	"wechatwebapi/models"
)

type Shakehandpubkey struct {
	V1PrivKey []byte
	V1pubKey  []byte
	V2PrivKey []byte
	V2pubKey  []byte
}

func (httpclient *HttpClientModel) InitMmtlsShake(ip, host string, PROXY models.ProxyInfo, Shakehandpubkey Shakehandpubkey) (*MmtlsClient, error) {
	httpclient.mmtlsClient.ClientSeq = 1
	httpclient.mmtlsClient.ServerSeq = 1
	var packetHeader MmtlsPacketHeader

	if len(Shakehandpubkey.V1pubKey) > 0 && len(Shakehandpubkey.V2pubKey) > 0 {
		httpclient.mmtlsClient.Shakehandprikey = Shakehandpubkey.V1PrivKey
		httpclient.mmtlsClient.Shakehandpubkey = Shakehandpubkey.V1pubKey

		httpclient.mmtlsClient.Shakehandprikey_2 = Shakehandpubkey.V2PrivKey
		httpclient.mmtlsClient.Shakehandpubkey_2 = Shakehandpubkey.V2pubKey
	} else {
		httpclient.mmtlsClient.Shakehandprikey, httpclient.mmtlsClient.Shakehandpubkey = Cilent.GenECDH415Key()
		httpclient.mmtlsClient.Shakehandprikey_2, httpclient.mmtlsClient.Shakehandpubkey_2 = Cilent.GenECDH415Key()

		if len(httpclient.mmtlsClient.Shakehandpubkey) <= 0 && len(httpclient.mmtlsClient.Shakehandpubkey_2) <= 0 {
			return &MmtlsClient{}, errors.New("Mmtls: GenECDH415Key 生成失败")
		}
	}

	httpclient.mmtlsClient.Shakehandpubkeylen = int32(len(httpclient.mmtlsClient.Shakehandpubkey))
	httpclient.mmtlsClient.Shakehandpubkeylen2 = int32(len(httpclient.mmtlsClient.Shakehandpubkey))

	var helloContent = new(bytes.Buffer)

	time := time.Now().Unix()

	helloContent.Write([]byte{0x01, 0x03, 0xf1, 0x01, 0xc0, 0x2b})
	helloContent.Write([]byte(Cilent.RandSeq(32)))
	binary.Write(helloContent, binary.BigEndian, (int32)(time))
	helloContent.Write([]byte{0x00, 0x00, 0x00, 0xa2, 0x01, 0x00, 0x00, 0x00, 0x9d, 0x00, 0x10, 0x02, 0x00, 0x00, 0x00, 0x47, 0x00, 0x00, 0x00, 0x01})
	binary.Write(helloContent, binary.BigEndian, int16(httpclient.mmtlsClient.Shakehandpubkeylen))
	helloContent.Write(httpclient.mmtlsClient.Shakehandpubkey[:httpclient.mmtlsClient.Shakehandpubkeylen])
	helloContent.Write([]byte{0x00, 0x00, 0x00, 0x47, 0x00, 0x00, 0x00, 0x02})
	binary.Write(helloContent, binary.BigEndian, int16(httpclient.mmtlsClient.Shakehandpubkeylen2))
	helloContent.Write(httpclient.mmtlsClient.Shakehandpubkey_2[:httpclient.mmtlsClient.Shakehandpubkeylen2])
	helloContent.Write([]byte{0x00, 0x00, 0x00, 0x01})

	var packhelloContent = new(bytes.Buffer)
	binary.Write(packhelloContent, binary.BigEndian, (int32)(helloContent.Len()))
	packhelloContent.Write(helloContent.Bytes())

	/*
		这里需要计算sha256
	*/
	mserverpubhashs := new(bytes.Buffer)
	mserverpubhashs.Write(packhelloContent.Bytes())

	var packsenddata = new(bytes.Buffer)
	packsenddata.Write([]byte{0x16, 0xf1, 0x03})
	binary.Write(packsenddata, binary.BigEndian, int16(packhelloContent.Len()))
	packsenddata.Write(packhelloContent.Bytes())

	uniquenumstr := "/mmtls/" + strconv.Itoa(int(time))

	recv_data, err := httpclient.POST(ip, uniquenumstr, packsenddata.Bytes(), host, PROXY)
	if err != nil {
		return &MmtlsClient{}, err
	}

	readerHeader := bytes.NewReader(recv_data[:5])
	mmtlsheader := []byte(string(recv_data[:5]))
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

	server_pubkeydata := recv_data[5 : 5+packetHeader.headerpacketLen]
	server_pubecdhkey := server_pubkeydata[58:len(server_pubkeydata)]
	//fmt.Println(hex.EncodeToString(server_pubecdhkey))
	mserverpubhashs.Write(server_pubkeydata)
	mserverpubhash_Bytes := Getsha256(mserverpubhashs.Bytes())

	httpclient.mmtlsClient.ShakehandECDHkey = Cilent.DoECDH415(httpclient.mmtlsClient.Shakehandprikey, server_pubecdhkey)
	if httpclient.mmtlsClient.ShakehandECDHkey == nil {
		return &MmtlsClient{}, errors.New("Mmtls: 秘钥交互失败")
	}
	if len(httpclient.mmtlsClient.ShakehandECDHkey) == 32 {
		httpclient.mmtlsClient.ShakehandECDHkeyLen = int32(len(httpclient.mmtlsClient.ShakehandECDHkey))
		m := sha256.New()
		m.Write(httpclient.mmtlsClient.ShakehandECDHkey[:])
		httpclient.mmtlsClient.ShakehandECDHkeyhash = m.Sum(nil)
		//shakehandECDHkey_string := hex.EncodeToString(tcpClient.mmtlsClient.shakehandECDHkeyhash)
		//fmt.Println(shakehandECDHkey_string)
		L := 56
		var infobytes = new(bytes.Buffer)
		infobytes.Write([]byte{0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e})
		infobytes.Write(mserverpubhash_Bytes)
		hkdfexpandkey := Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.ShakehandECDHkeyhash, infobytes.Bytes(), L)
		//fmt.Println(len(hkdfexpandkey))
		httpclient.mmtlsClient.Encrptmmtlskey = hkdfexpandkey[:16]
		httpclient.mmtlsClient.Decryptmmtlskey = hkdfexpandkey[16:32]
		httpclient.mmtlsClient.EncrptmmtlsIv = hkdfexpandkey[32:44]
		httpclient.mmtlsClient.DecryptmmtlsIv = hkdfexpandkey[44:56]

		var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.DecryptmmtlsIv[8:12])
		var xorkeyint uint32
		binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
		xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ServerSeq)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, xorkeyint)

		var decryptmmtlsIv_seq = new(bytes.Buffer)
		decryptmmtlsIv_seq.Write(httpclient.mmtlsClient.DecryptmmtlsIv[:8])
		decryptmmtlsIv_seq.Write(buf.Bytes())

		httpclient.mmtlsClient.CurDecryptSeqIv = decryptmmtlsIv_seq.Bytes()
		httpclient.mmtlsClient.ServerSeq += 1

		datapos := 5 + packetHeader.headerpacketLen
		recv_data = recv_data[datapos:len(recv_data)]

		readerHeader := bytes.NewReader(recv_data[:5])
		mmtlsheader = []byte(string(recv_data[0:5]))
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

		mmtls_part_2_data := recv_data[5 : 5+packetHeader.headerpacketLen]
		var mmtls_part_2_aad = new(bytes.Buffer)
		mmtls_part_2_aad.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
		mmtls_part_2_aad.Write(mmtlsheader)
		/*
			aes_decrypt
		*/
		decrypt_part2 := Cilent.NewAES_GCMDecrypter(httpclient.mmtlsClient.Decryptmmtlskey, mmtls_part_2_data, httpclient.mmtlsClient.CurDecryptSeqIv, mmtls_part_2_aad.Bytes())
		if decrypt_part2 != nil {
			mserverpubhashs.Write(decrypt_part2[:len(decrypt_part2)])
			httpclient.mmtlsClient.Decrypt_part2_hash256 = Getsha256(mserverpubhashs.Bytes())

			datapos := 5 + packetHeader.headerpacketLen
			recv_data = recv_data[datapos:len(recv_data)]
			readerHeader := bytes.NewReader(recv_data[:5])
			mmtlsheader = []byte(string(recv_data[0:5]))
			binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
			binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
			binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

			mmtls_part_3_data := recv_data[5 : 5+packetHeader.headerpacketLen]
			mmtls_part_3_aad := new(bytes.Buffer)
			mmtls_part_3_aad.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
			mmtls_part_3_aad.Write(mmtlsheader)

			var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.DecryptmmtlsIv[8:12])
			var xorkeyint uint32
			binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
			xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ServerSeq)
			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, xorkeyint)

			var decryptmmtlsIv_seq = new(bytes.Buffer)
			decryptmmtlsIv_seq.Write(httpclient.mmtlsClient.DecryptmmtlsIv[:8])
			decryptmmtlsIv_seq.Write(buf.Bytes())

			httpclient.mmtlsClient.CurDecryptSeqIv = decryptmmtlsIv_seq.Bytes()

			decrypt_part3 := Cilent.NewAES_GCMDecrypter(httpclient.mmtlsClient.Decryptmmtlskey, mmtls_part_3_data, httpclient.mmtlsClient.CurDecryptSeqIv, mmtls_part_3_aad.Bytes())

			if decrypt_part3 != nil {
				//fmt.Println(hex.EncodeToString(decrypt_part3))
				httpclient.mmtlsClient.ServerSeq += 1
				mserverpubhashs.Write(decrypt_part3[:len(decrypt_part3)])
				httpclient.mmtlsClient.Decrypt_part3_hash256 = Getsha256(mserverpubhashs.Bytes())

				var HkdfExpand_info_access = new(bytes.Buffer)
				HkdfExpand_info_access.Write([]byte{0x50, 0x53, 0x4b, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53})
				HkdfExpand_info_access.Write(httpclient.mmtlsClient.Decrypt_part2_hash256)
				httpclient.mmtlsClient.Hkdfexpand_pskaccess_key = Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.ShakehandECDHkeyhash, HkdfExpand_info_access.Bytes(), 32)

				var HkdfExpand_info_refresh = new(bytes.Buffer)
				HkdfExpand_info_refresh.Write([]byte{0x50, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48})
				HkdfExpand_info_refresh.Write(httpclient.mmtlsClient.Decrypt_part2_hash256)
				httpclient.mmtlsClient.Hkdfexpand_pskrefresh_key = Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.ShakehandECDHkeyhash, HkdfExpand_info_refresh.Bytes(), 32)

				datapos := 5 + packetHeader.headerpacketLen
				recv_data = recv_data[datapos:len(recv_data)]
				readerHeader := bytes.NewReader(recv_data[:5])
				mmtlsheader = []byte(string(recv_data[0:5]))
				binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
				binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
				binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

				mmtls_part_4_data := recv_data[5 : 5+packetHeader.headerpacketLen]
				mmtls_part_4_aad := new(bytes.Buffer)
				mmtls_part_4_aad.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03})
				mmtls_part_4_aad.Write(mmtlsheader)

				var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.DecryptmmtlsIv[8:12])
				var xorkeyint uint32
				binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
				xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ServerSeq)
				buf := new(bytes.Buffer)
				binary.Write(buf, binary.BigEndian, xorkeyint)

				var decryptmmtlsIv_seq = new(bytes.Buffer)
				decryptmmtlsIv_seq.Write(httpclient.mmtlsClient.DecryptmmtlsIv[:8])
				decryptmmtlsIv_seq.Write(buf.Bytes())
				httpclient.mmtlsClient.ServerSeq += 1
				httpclient.mmtlsClient.CurDecryptSeqIv = decryptmmtlsIv_seq.Bytes()

				decrypt_part4 := Cilent.NewAES_GCMDecrypter(httpclient.mmtlsClient.Decryptmmtlskey, mmtls_part_4_data, httpclient.mmtlsClient.CurDecryptSeqIv, mmtls_part_4_aad.Bytes())
				if decrypt_part4 != nil {
					//log.Println(hex.EncodeToString(decrypt_part4))
					var HkdfExpand_info_serverfinish = new(bytes.Buffer)
					HkdfExpand_info_serverfinish.Write([]byte{0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64})
					httpclient.mmtlsClient.HkdfExpand_info_serverfinish_key = Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.ShakehandECDHkeyhash, HkdfExpand_info_serverfinish.Bytes(), 32)
					sh := hmac.New(sha256.New, httpclient.mmtlsClient.HkdfExpand_info_serverfinish_key)
					sh.Write(httpclient.mmtlsClient.Decrypt_part3_hash256)
					//serverdigest_m := sh.Sum(nil)
					//log.Println(hex.EncodeToString(serverdigest_m))

					//拿下earlydatapar3
					var earlydatalen int32
					earlydata := bytes.NewReader(decrypt_part3[6:10])
					binary.Read(earlydata, binary.BigEndian, &earlydatalen)
					httpclient.mmtlsClient.Earlydatapart = decrypt_part3[6 : 6+4+earlydatalen]
					return httpclient.mmtlsClient, nil
				}
			}
		}
		return &MmtlsClient{}, errors.New("Mmtls: ASEDecrypt 解密失败")

	}
	return &MmtlsClient{}, errors.New("Mmtls: 交互秘钥长度存在异常")
}
