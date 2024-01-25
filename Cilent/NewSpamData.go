package Cilent

/*
#include "encrypt.h"
#cgo !windows LDFLAGS: -L.
*/
import "C"

import (
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"hash/crc32"
	"time"
	"unsafe"
	"wechatwebapi/Cilent/device"
	"wechatwebapi/Cilent/mm"
)

func GetNewSpamData(Deviceid, DeviceName string) []byte {
	datas, _ := hex.DecodeString(SAEDATSTRING)
	wx_saedata_03 := DoZlibUnCompress(datas)
	saedatapbdata := &mm.Wcaes{}
	err := proto.Unmarshal(wx_saedata_03, saedatapbdata)
	if err == nil {
		SaeIv := saedatapbdata.IV
		SaeTable := saedatapbdata.Tablekey
		SaeValue := saedatapbdata.Tablevalue
		SaeFinalData := saedatapbdata.Unkown18
		timeStamp := int(time.Now().Unix())
		xorKey := uint8((timeStamp * 0xffffffed) + 7)

		uuid1, uuid2 := device.Uuid(Deviceid)


		if len(Deviceid) < 32 {
			Dlen := 32 - len(Deviceid)
			Fill := "ff95DODUJ4EysYiogKZSmajWCUKUg9RX"
			Deviceid = Deviceid + Fill[:Dlen]
		}

		spamDataBody := &mm.SpamDataBody{
			UnKnown1:              proto.Int32(1),
			TimeStamp:             proto.Int32(int32(timeStamp)),
			KeyHash:               proto.Int32(int32(C.makeKeyHash(C.int(xorKey)))),
			Yes1:                  proto.String(XorEncodeStr("yes", xorKey)),
			Yes2:                  proto.String(XorEncodeStr("yes", xorKey)),
			IosVersion:            proto.String(XorEncodeStr("13.3.1", xorKey)),
			DeviceType:            proto.String(XorEncodeStr("iPad", xorKey)),
			UnKnown2:              proto.Int32(2),
			IdentifierForVendor:   proto.String(XorEncodeStr(uuid1, xorKey)),
			AdvertisingIdentifier: proto.String(XorEncodeStr(uuid2, xorKey)),
			Carrier:               proto.String(XorEncodeStr("中国移动", xorKey)),
			BatteryInfo:           proto.Int32(1),
			NetworkName:           proto.String(XorEncodeStr("en0", xorKey)),
			NetType:               proto.Int32(1),
			AppBundleId:           proto.String(XorEncodeStr("com.tencent.xin", xorKey)),
			DeviceName:            proto.String(XorEncodeStr(DeviceName, xorKey)),
			UserName:              proto.String(XorEncodeStr("iPad11,3", xorKey)),
			Unknown3:              proto.Int64(device.DeviceNumber(Deviceid[:29] + "FFF")),
			Unknown4:              proto.Int64(device.DeviceNumber(Deviceid[:29] + "OOO")),
			Unknown5:              proto.Int32(1),
			Unknown6:              proto.Int32(4),
			Lang:                  proto.String(XorEncodeStr("zh", xorKey)),
			Country:               proto.String(XorEncodeStr("CN", xorKey)),
			Unknown7:              proto.Int32(4),
			DocumentDir:           proto.String(XorEncodeStr("/var/mobile/Containers/Data/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x10101201))+"/Documents", xorKey)),
			Unknown8:              proto.Int32(0),
			Unknown9:              proto.Int32(0),
			HeadMD5:               proto.String(XorEncodeStr(device.GetCidMd5(Deviceid,device.GetCid(0x0262626262626)), xorKey)),
			AppUUID:               proto.String(XorEncodeStr(uuid1, xorKey)),
			SyslogUUID:            proto.String(XorEncodeStr(uuid2, xorKey)),
			Unknown10:             proto.String(""),
			Unknown11:             proto.String(""),
			AppName:               proto.String(XorEncodeStr("微信", xorKey)),
			SshPath:               proto.String(""),
			TempTest:              proto.String(""),
			DevMD5:                proto.String(""),
			DevUser:               proto.String(""),
			Unknown12:             proto.String(""),
			IsModify:              proto.Int32(0),
			ModifyMD5:             proto.String(""),
			RqtHash:               proto.Int64(0x0555555555555555),
		}
		wxFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/WeChat", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000001)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, wxFile)

		opensslFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/OpenSSL.framework/OpenSSL", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000002)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, opensslFile)

		protoFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/ProtobufLite.framework/ProtobufLite", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000003)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, protoFile)

		marsbridgenetworkFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/marsbridgenetwork.framework/marsbridgenetwork", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000004)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, marsbridgenetworkFile)

		matrixreportFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/matrixreport.framework/matrixreport", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000005)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, matrixreportFile)

		andromedaFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/andromeda.framework/andromeda", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000006)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, andromedaFile)

		marsFile := &mm.FileInfo{
			Fileuuid: proto.String(XorEncodeStr("/var/containers/Bundle/Application/"+device.GetCidUUid(Deviceid,device.GetCid(0x098521236654))+"/WeChat.app/Frameworks/mars.framework/mars", xorKey)),
			Filepath: proto.String(XorEncodeStr(device.GetCidUUid(Deviceid,device.GetCid(0x30000007)), xorKey)),
		}
		spamDataBody.AppFileInfo = append(spamDataBody.AppFileInfo, marsFile)
		srcdata, _ := proto.Marshal(spamDataBody)

		newClientCheckData := &mm.NewClientCheckData{
			C32Cdata:  proto.Int64(int64(crc32.ChecksumIEEE([]byte(srcdata)))),
			TimeStamp: proto.Int64(int64(time.Now().Unix())),
			Databody:  srcdata,
		}

		ccddata, _ := proto.Marshal(newClientCheckData)
		//需要加密的数据
		compressdata := DoZlibCompress([]byte(ccddata))
		bytesSaeIv := make([]byte, len(SaeIv))
		bytesSaeIv = SaeIv[:len(SaeIv)]
		var compressdata_ptr = (*C.char)(unsafe.Pointer(&compressdata[0]))
		var SaeIv_ptr = (*C.char)(unsafe.Pointer(&bytesSaeIv[0]))
		var SaeTable_ptr = (*C.char)(unsafe.Pointer(&SaeTable[0]))
		var SaeValue_ptr = (*C.char)(unsafe.Pointer(&SaeValue[0]))
		var SaeFinalData_ptr = (*C.char)(unsafe.Pointer(&SaeFinalData[0]))
		var outbuffer = make([]byte, len(compressdata)*2)
		var outbuffer_ptr = (*C.char)(unsafe.Pointer(&outbuffer[0]))
		var outbuffer_len uint32
		var outbuffer_len_ptr = (*C.uint)(unsafe.Pointer(&outbuffer_len))
		cs := C.CString("")
		ret := C.nativewcswbaes(cs, compressdata_ptr, C.uint(len(compressdata)), SaeIv_ptr,
			C.uint(len(bytesSaeIv)), SaeTable_ptr, C.uint(len(SaeTable)), SaeValue_ptr, C.uint(len(SaeValue)),
			SaeFinalData_ptr, C.uint(len(SaeFinalData)), 0x3060, outbuffer_ptr, outbuffer_len_ptr)
		if ret == 1 {
			return outbuffer[:outbuffer_len]
		}
	}
	return nil
}
