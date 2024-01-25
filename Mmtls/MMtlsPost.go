package Mmtls

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"time"
	"wechatwebapi/models"
)

// mmtlspost
func (httpclient *HttpClientModel) MMtlsPost(ip, host, cgiurl string, data []byte, P models.ProxyInfo) ([]byte, error) {
	var err error
	newsenddata := new(bytes.Buffer)
	binary.Write(newsenddata, binary.BigEndian, int16(len(cgiurl)))
	newsenddata.Write([]byte(cgiurl))
	binary.Write(newsenddata, binary.BigEndian, int16(len(host)))
	newsenddata.Write([]byte(host))
	binary.Write(newsenddata, binary.BigEndian, int32(len(data)))
	newsenddata.Write(data)
	send_data := new(bytes.Buffer)
	binary.Write(send_data, binary.BigEndian, int32(newsenddata.Len()))
	send_data.Write(newsenddata.Bytes())
	encryptdata := httpclient.MmtlsEncryptData(send_data.Bytes())
	if encryptdata == nil {
		return []byte{}, errors.New("MMTLS: 数据[EncryptData]失败")
	}
	var recv_data []byte

	uniquenumstr := "/mmtls/" + strconv.Itoa(int(time.Now().Unix()))

	recv_data, err = httpclient.POST(ip, uniquenumstr, encryptdata, host, P)

	if err != nil {
		return []byte{}, err
	}

	response := new(bytes.Buffer)
	/*Separate := Separate(recv_data)
	for _, v := range Separate {
		response.Write(httpclient.MmtlsDecryptData(v))
	}*/

	response.Write(httpclient.MmtlsDecryptData(recv_data))

	if response.Bytes() == nil {
		return []byte{}, errors.New("MMTLS: 数据[DecryptData]失败")
	}

	return response.Bytes(), nil
}
