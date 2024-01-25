package Mmtls

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
	"wechatwebapi/Cilent"
)

func (httpclient *HttpClientModel) MmtlsEncryptData(srcbindata []byte) []byte {
	httpclient.mmtlsClient.ClientSeq = 1
	httpclient.mmtlsClient.ServerSeq = 1

	var newsendbuffer = new(bytes.Buffer)
	newsendbuffer.Write([]byte{0x01, 0x03, 0xf1, 0x01, 0x00, 0xa8})
	newsendbuffer.Write([]byte(Cilent.RandSeq(32)))
	time := time.Now().Unix()
	binary.Write(newsendbuffer, binary.BigEndian, (int32)(time))
	newsendbuffer.Write([]byte{0x00, 0x00, 0x00, 0x6f, 0x01, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x0f, 0x01})
	newsendbuffer.Write(httpclient.mmtlsClient.Earlydatapart)
	var newsendbuffer_content = new(bytes.Buffer)

	binary.Write(newsendbuffer_content, binary.BigEndian, int32(newsendbuffer.Len()))
	newsendbuffer_content.Write(newsendbuffer.Bytes())

	//自定义newsendbufferhash
	/*newsendbufferhash := sha256.New()
	newsendbufferhash.Write(newsendbuffer_content.Bytes())
	newsendbufferhashhash256 := newsendbufferhash.Sum(nil)*/

	newsendbufferhashs := new(bytes.Buffer)
	newsendbufferhashs.Write(newsendbuffer_content.Bytes())
	newsendbufferhashhash256 := Getsha256(newsendbufferhashs.Bytes())

	HkdfExpand_early_data := new(bytes.Buffer)
	HkdfExpand_early_data.Write([]byte{0x65, 0x61, 0x72, 0x6c, 0x79, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e})
	HkdfExpand_early_data.Write(newsendbufferhashhash256)

	hkdfexpand_shortencryptkey := Cilent.Hkdf_Expand(sha256.New, httpclient.mmtlsClient.Hkdfexpand_pskaccess_key, HkdfExpand_early_data.Bytes(), 28)
	httpclient.mmtlsClient.Encrptshortmmtlskey = hkdfexpand_shortencryptkey[:16]
	httpclient.mmtlsClient.Encrptshortmmtlsiv = hkdfexpand_shortencryptkey[16:28]

	part1_data := new(bytes.Buffer)
	part1_data.Write([]byte{0x19, 0xf1, 0x03})
	binary.Write(part1_data, binary.BigEndian, int16(len(newsendbuffer_content.Bytes())))
	part1_data.Write(newsendbuffer_content.Bytes())

	//业务数据
	part2_data_inputdata := new(bytes.Buffer)
	part2_data_inputdata.Write([]byte{0x00, 0x00, 0x00, 0x10, 0x08, 0x00, 0x00, 0x00, 0x0b, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00, 0x12})
	binary.Write(part2_data_inputdata, binary.BigEndian, int32(int(time)))

	var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.Encrptshortmmtlsiv[8:12])
	var xorkeyint uint32
	binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
	xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ClientSeq)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, xorkeyint)

	var encryptmmtlsIv_seq = new(bytes.Buffer)
	encryptmmtlsIv_seq.Write(httpclient.mmtlsClient.Encrptshortmmtlsiv[:8])
	encryptmmtlsIv_seq.Write(buf.Bytes())

	var aad_noop = new(bytes.Buffer)

	binary.Write(aad_noop, binary.BigEndian, uint64(httpclient.mmtlsClient.ClientSeq))
	aad_noop.Write([]byte{0x19, 0xf1, 0x03})
	binary.Write(aad_noop, binary.BigEndian, int16(part2_data_inputdata.Len()+0x10))

	aes_gsm_data := Cilent.NewAES_GCMEncrypter(httpclient.mmtlsClient.Encrptshortmmtlskey, part2_data_inputdata.Bytes(), encryptmmtlsIv_seq.Bytes(), aad_noop.Bytes())
	if aes_gsm_data != nil {
		httpclient.mmtlsClient.ClientSeq += 1
		part2_data := new(bytes.Buffer)
		part2_data.Write([]byte{0x19, 0xf1, 0x03})
		binary.Write(part2_data, binary.BigEndian, int16(part2_data_inputdata.Len()+0x10))
		part2_data.Write(aes_gsm_data)

		part3_data_inputdata := srcbindata

		var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.Encrptshortmmtlsiv[8:12])
		var xorkeyint uint32
		binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
		xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ClientSeq)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, xorkeyint)

		var encryptmmtlsIv_seq = new(bytes.Buffer)
		encryptmmtlsIv_seq.Write(httpclient.mmtlsClient.Encrptshortmmtlsiv[:8])
		encryptmmtlsIv_seq.Write(buf.Bytes())

		var aad_noop = new(bytes.Buffer)
		binary.Write(aad_noop, binary.BigEndian, uint64(httpclient.mmtlsClient.ClientSeq))
		aad_noop.Write([]byte{0x17, 0xf1, 0x03})
		binary.Write(aad_noop, binary.BigEndian, int16(len(part3_data_inputdata)+0x10))

		aes_gsm_data := Cilent.NewAES_GCMEncrypter(httpclient.mmtlsClient.Encrptshortmmtlskey, part3_data_inputdata, encryptmmtlsIv_seq.Bytes(), aad_noop.Bytes())
		if aes_gsm_data != nil {
			httpclient.mmtlsClient.ClientSeq += 1
			part3_data := new(bytes.Buffer)
			part3_data.Write([]byte{0x17, 0xf1, 0x03})
			binary.Write(part3_data, binary.BigEndian, int16(len(part3_data_inputdata)+0x10))
			part3_data.Write(aes_gsm_data)

			part4_data_inputdata := new(bytes.Buffer)
			part4_data_inputdata.Write([]byte{0x00, 0x00, 0x00, 0x03, 0x00, 0x01, 0x01})

			var xorkeyBuffer = bytes.NewReader(httpclient.mmtlsClient.Encrptshortmmtlsiv[8:12])
			var xorkeyint uint32
			binary.Read(xorkeyBuffer, binary.BigEndian, &xorkeyint)
			xorkeyint = xorkeyint ^ uint32(httpclient.mmtlsClient.ClientSeq)
			buf := new(bytes.Buffer)
			binary.Write(buf, binary.BigEndian, xorkeyint)

			var encryptmmtlsIv_seq = new(bytes.Buffer)
			encryptmmtlsIv_seq.Write(httpclient.mmtlsClient.Encrptshortmmtlsiv[:8])
			encryptmmtlsIv_seq.Write(buf.Bytes())

			var aad_noop = new(bytes.Buffer)
			binary.Write(aad_noop, binary.BigEndian, uint64(httpclient.mmtlsClient.ClientSeq))
			aad_noop.Write([]byte{0x15, 0xf1, 0x03})
			binary.Write(aad_noop, binary.BigEndian, int16(len(part4_data_inputdata.Bytes())+0x10))

			aes_gsm_data := Cilent.NewAES_GCMEncrypter(httpclient.mmtlsClient.Encrptshortmmtlskey, part4_data_inputdata.Bytes(), encryptmmtlsIv_seq.Bytes(), aad_noop.Bytes())
			if aes_gsm_data != nil {
				httpclient.mmtlsClient.ClientSeq += 1
				part4_data := new(bytes.Buffer)
				part4_data.Write([]byte{0x15, 0xf1, 0x03})
				binary.Write(part4_data, binary.BigEndian, int16(len(part4_data_inputdata.Bytes())+0x10))
				part4_data.Write(aes_gsm_data)

				finalSendData := new(bytes.Buffer)
				finalSendData.Write(part1_data.Bytes())
				//log.Println(hex.EncodeToString(part1_data.Bytes()))
				finalSendData.Write(part2_data.Bytes())
				//log.Println(hex.EncodeToString(part2_data.Bytes()))
				finalSendData.Write(part3_data.Bytes())
				//log.Println(hex.EncodeToString(part3_data.Bytes()))
				finalSendData.Write(part4_data.Bytes())
				//log.Println(hex.EncodeToString(part4_data.Bytes()))
				newsendbufferhashs.Write(part2_data_inputdata.Bytes())
				httpclient.mmtlsClient.Newsendbufferhashs = newsendbufferhashs.Bytes()
				return finalSendData.Bytes()
			}
		}
	}
	return nil
}
