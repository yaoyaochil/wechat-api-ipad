package Mmtls

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"wechatwebapi/Cilent"
)

func (httpclient *HttpClientModel) MmtlsDecryptData(srcbindata []byte) []byte {

	var packetHeader MmtlsPacketHeader
	recv_data := srcbindata
	readerHeader := bytes.NewReader(recv_data[:5])
	mmtlsheader := []byte(string(recv_data[:5]))
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

	server_data := recv_data[5 : 5+packetHeader.headerpacketLen]
	newsendbufferhashs := new(bytes.Buffer)
	newsendbufferhashs.Write(httpclient.mmtlsClient.Newsendbufferhashs)
	newsendbufferhashs.Write(server_data)
	decrypt_serverdata_hash256 := Getsha256(newsendbufferhashs.Bytes())

	var HkdfExpand_handshake = new(bytes.Buffer)
	HkdfExpand_handshake.Write([]byte{0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e})
	HkdfExpand_handshake.Write(decrypt_serverdata_hash256)

	HkdfExpand_handshake_key := Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.Hkdfexpand_pskaccess_key, HkdfExpand_handshake.Bytes(), 28)
	httpclient.mmtlsClient.Decrptshortmmtlskey = HkdfExpand_handshake_key[:16]
	httpclient.mmtlsClient.Decrptshortmmtlsiv = HkdfExpand_handshake_key[16:28]

	datapos := 5 + packetHeader.headerpacketLen

	recv_data = recv_data[datapos:len(recv_data)]
	readerHeader = bytes.NewReader(recv_data[:5])
	mmtlsheader = []byte(string(recv_data[:5]))
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
	binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

	serverfinishdata := recv_data[5 : 5+packetHeader.headerpacketLen]
	serverfinishdata_aad := new(bytes.Buffer)
	serverfinishdata_aad.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	serverfinishdata_aad.Write(mmtlsheader)

	var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.Decrptshortmmtlsiv[8:12])
	var xorkeyint uint32
	binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
	xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ServerSeq)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, xorkeyint)

	var decryptmmtlsIv_seq = new(bytes.Buffer)
	decryptmmtlsIv_seq.Write(httpclient.mmtlsClient.Decrptshortmmtlsiv[:8])
	decryptmmtlsIv_seq.Write(buf.Bytes())

	decryptserverfinishdata := Cilent.NewAES_GCMDecrypter(httpclient.mmtlsClient.Decrptshortmmtlskey, serverfinishdata, decryptmmtlsIv_seq.Bytes(), serverfinishdata_aad.Bytes())
	if decryptserverfinishdata != nil {
		httpclient.mmtlsClient.ServerSeq += 1
		var HkdfExpand_serverfinish_datakey = new(bytes.Buffer)
		HkdfExpand_serverfinish_datakey.Write([]byte{0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64})
		hkdfexpand_shortencryptkey := Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.Hkdfexpand_pskaccess_key, HkdfExpand_serverfinish_datakey.Bytes(), 32)

		h := hmac.New(sha256.New, hkdfexpand_shortencryptkey)
		h.Write(decrypt_serverdata_hash256)
		//clientdigest_m := h.Sum(nil)

		datapos := 5 + packetHeader.headerpacketLen

		recv_data = recv_data[datapos:len(recv_data)]
		readerHeader = bytes.NewReader(recv_data[:5])
		mmtlsheader = []byte(string(recv_data[:5]))
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerbyte)
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerversion)
		binary.Read(readerHeader, binary.BigEndian, &packetHeader.headerpacketLen)

		businessdata := recv_data[5 : 5+packetHeader.headerpacketLen]
		businessdata_aad := new(bytes.Buffer)
		businessdata_aad.Write([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
		businessdata_aad.Write(mmtlsheader)

		var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.Decrptshortmmtlsiv[8:12])
		var xorkeyint uint32
		binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
		xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ServerSeq)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, xorkeyint)

		var decryptmmtlsIv_seq = new(bytes.Buffer)
		decryptmmtlsIv_seq.Write(httpclient.mmtlsClient.Decrptshortmmtlsiv[:8])
		decryptmmtlsIv_seq.Write(buf.Bytes())

		decrypt_businessdata := Cilent.NewAES_GCMDecrypter(httpclient.mmtlsClient.Decrptshortmmtlskey, businessdata, decryptmmtlsIv_seq.Bytes(), businessdata_aad.Bytes())
		if decrypt_businessdata != nil {
			return decrypt_businessdata
		}

	}
	return nil
}
