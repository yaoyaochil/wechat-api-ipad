package device

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"wechatwebapi/Fun"
)

func Imei(DeviceId string) string {
	return DeviceId
}

func SoftType(DeviceId string) string {
	uuid1, uuid2 := Uuid(DeviceId)
	return "<softtype><k3>13.3.1</k3><k9>iPad</k9><k10>2</k10><k19>" + uuid1 + "</k19><k20>" + uuid2 + "</k20><k21>iPad5</k21><k22>(null)</k22><k24>" + Mac(DeviceId) + "</k24><k33>\\345\\276\\256\\344\\277\\241</k33><k47>1</k47><k50>0</k50><k51>com.tencent.xin</k51><k54>iPad11,3</k54><k61>2</k61></softtype>"
}

func Uuid(DeviceId string) (uuid1 string, uuid2 string) {
	Md5DataA := MD5ToLower(DeviceId + "SM202003220432")
	uuid1 = fmt.Sprintf("%x-%x-%x-%x-%x",Md5DataA[0:8],Md5DataA[2:6],Md5DataA[3:7],Md5DataA[1:5],Md5DataA[20:32])
	Md5DataB := MD5ToLower(DeviceId + "BM202003220432")
	uuid2 = fmt.Sprintf("%x-%x-%x-%x-%x",Md5DataB[0:8],Md5DataB[2:6],Md5DataB[3:7],Md5DataB[1:5],Md5DataB[20:32])
	return
}

func Mac(DeviceId string) string {
	Md5Data := MD5ToLower(DeviceId + "CP202003220432")
	return fmt.Sprintf("3C:2E:F9:%v:%v:%v",Md5Data[5:7],Md5Data[7:9],Md5Data[10:12])
}

func GetCid(s int) string {
	M := inttobytes(s >> 12)
	return hex.EncodeToString(M)
}

func GetCidUUid(DeviceId, Cid string) string {
	Md5Data := MD5ToLower(DeviceId + Cid)
	return fmt.Sprintf("%x-%x-%x-%x-%x",Md5Data[0:8],Md5Data[2:6],Md5Data[3:7],Md5Data[1:5],Md5Data[20:32])
}


func GetCidMd5(DeviceId, Cid string) string {
	Md5Data := MD5ToLower(DeviceId + Cid)
	return "A136" + Md5Data[5:]
}

func inttobytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func DeviceNumber(DeviceId string) int64 {
	ssss := []byte(MD5ToLower(DeviceId))
	ccc := Fun.Hex2int(&ssss) >> 8
	ddd := ccc + 60000000000000000
	if ddd > 80000000000000000 {
		ddd = ddd - (80000000000000000 - ddd)
	}
	return int64(ddd)
}

