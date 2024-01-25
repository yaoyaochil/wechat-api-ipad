package Login

var version = 0x27000d34
var deviceType_byte = []byte("android-27")
var deviceType_str = "android-27"

//func A16Login(Data ManualAuthReq, IP, domain string) wxCilent.ResponseResult {
//	if Data.UserName == "" || Data.Password == "" {
//		return wxCilent.ResponseResult{
//			Code:    -8,
//			Success: false,
//			Message: "请输入账号或密码",
//			Data:    nil,
//		}
//	}
//
//	if Data.DeviceData != "" {
//		Data.DeviceData = Data.DeviceData[0:15]
//	}
//
//	//初始化Mmtls
//	httpclient, MmtlsClient, err := ioscomm.MmtlsInitialize(Data.Proxy)
//	if err != nil {
//		return wxCilent.ResponseResult{
//			Code:    -8,
//			Success: false,
//			Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
//			Data:    nil,
//		}
//	}
//
//	ai := &appproto.AccountInfo{
//
//		Username:      Data.UserName,
//		Password:      Data.Password,
//		Deviceid:      []byte(Data.DeviceData),
//		Deviceid_str:  Data.DeviceData,
//		Sessionkey:    []byte{},
//		Clientversion: 0x27000d34,
//		Devicetype:    "android-27",
//		Devicetoken:   "",
//
//		//softtype
//		AllowLocation:            0,
//		IsDebug:                  0,
//		IsRoot:                   0,
//		RadioVerion:              "M8994F-2.6.42.5.03",
//		RoBuildVersionRelease:    "8.1.0",
//		Imei:                     comm.AndriodImei(Data.DeviceData),
//		AndriodID:                comm.AndriodID(Data.DeviceData),
//		SerialID:                 comm.AndriodSerial(Data.DeviceData),
//		AndroidOsBuildModel:      "Nexus 5X",
//		CPUCount:                 6,
//		Hardware:                 "Qualcomm Technologies, Inc MSM8992",
//		Revision:                 "",
//		Serial:                   "",
//		Ssid:                     "<unknown ssid>",
//		Bssid:                    comm.AndriodBssid(Data.DeviceData),
//		Features:                 "half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt evtstrm aes pmull sha1 sha2 crc32",
//		PackageSign:              comm.AndriodPackageSign(Data.DeviceData),
//		WifiName:                 "Chinanet-2.4G-103",
//		WifiFullName:             "&quot;Chinanet-2.4G-103&quot;",
//		FingerPrint:              "google/bullhead/bullhead:8.1.0/OPM7.181105.004/5038062:user/release-keys",
//		AndroidOsBuildBoard:      "bullhead",
//		AndroidOsBuildBootLoader: "BHZ32c",
//		AndroidOsBuildBRAND:      "google",
//		AndroidOsBuildDEVICE:     "bullhead",
//		AndroidOsBuildHARDWARE:   "bullhead",
//		AndroidOsBuildPRODUCT:    "bullhead",
//		RoProductManufacturer:    "LGE",
//		PhoneNumber:              "",
//		NetType:                  "wifi",
//		RecentTasks:              1,
//		PackageBuildNumber:       1640,
//		XMLInfoType:              3,
//		FeatureID:                "",
//		SoterID:                  "",
//		OAID:                     "",
//		IsRePack:                 0,
//		DataDirectory:            "data/user/0/com.tencent.mm/",
//		PackageName:              "com.tencent.mm",
//		IsQemu:                   0,
//		SimOperatorName:          "",
//		CashDBOpenSuccess:        "",
//		CPUDescription:           "0 ",
//		SubscriberID:             "",
//		SimSerialNumber:          "",
//		BlueToothAddress:         "",
//		KernelReleaseNumber:      "3.10.73-g0a05126d69c9",
//		WLanAddress:              comm.AndriodWLanAddress(Data.DeviceData),
//		Arch:                     "armeabi-v7a",
//	}
//
//	tii := &appproto.TrustInfoInit{}
//	tii.SetAccountInfo(ai)
//	tiiData := tii.ToBuffer()
//	hec := &comm.HybridEcdhClient{}
//	hec.Init()
//	hecData := hec.Encrypt(tiiData)
//	hypack := comm.PackHybridEcdh(tii.GetCmdid(), 10002, 0, nil, hecData)
//
//	recvData, err := httpclient.MMtlsPost(IP, domain, tii.GetUri(), hypack, Data.Proxy)
//
//	if err != nil {
//		return wxCilent.ResponseResult{
//			Code:    -8,
//			Success: false,
//			Message: fmt.Sprintf("系统异常：%v", err.Error()),
//			Data:    nil,
//		}
//	}
//
//	ph := comm.UnpackHybridEcdh(recvData)
//	devicetoken := hec.Decrypt(ph.Data)
//
//	tii.OnResponse(devicetoken)
//
//	secauth := &appproto.SecManualAuth{}
//	secauth.SetAccountInfo(ai)
//
//	tiData := secauth.ToBuffer()
//
//	hec1 := &comm.HybridEcdhClient{}
//	hec1.Init()
//	hecData1 := hec1.Encrypt(tiData)
//
//	hypack1 := comm.PackHybridEcdh(secauth.GetCmdId(), 10002, ai.UiCryptin, nil, hecData1)
//
//	recvData1, err := httpclient.MMtlsPost(IP, domain, secauth.GetUri(), hypack1, Data.Proxy)
//
//	if err != nil {
//		return wxCilent.ResponseResult{
//			Code:    -8,
//			Success: false,
//			Message: fmt.Sprintf("系统异常：%v", err.Error()),
//			Data:    nil,
//		}
//	}
//
//	ph1 := comm.UnpackHybridEcdh(recvData1)
//
//	ai.Cookies = ph1.Cookies
//	ai.UiCryptin = ph1.UICrypt
//
//	//secauth.OnResponse(hec1.Decrypt(ph1.Data))
//
//	loginRes := mmproto.UnifyAuthResponse{}
//	err = proto.Unmarshal(hec1.Decrypt(ph1.Data), &loginRes)
//
//	if err != nil {
//		return wxCilent.ResponseResult{
//			Code:    -8,
//			Success: false,
//			Message: fmt.Sprintf("解包失败：%v", err.Error()),
//			Data:    nil,
//		}
//	}
//
//	if loginRes.GetBaseResponse().GetRet() == 0 {
//		var LoginData ioscomm.LoginData
//		LoginData.Cooike = ph1.Cookies
//		LoginData.Mmtlsip = IP
//		LoginData.MmtlsHost = domain
//		LoginData.Deviceid_str = Data.DeviceData
//		LoginData.Deviceid_byte = []byte(Data.DeviceData)
//		LoginData.MmtlsKey = MmtlsClient
//		LoginData.ClientVersion = version
//		LoginData.DeviceType = deviceType_str
//		//保存redis
//		err = secauth.OnResponse(loginRes, LoginData)
//
//		if err != nil {
//			return wxCilent.ResponseResult{
//				Code:    -8,
//				Success: false,
//				Message: fmt.Sprintf("系统异常：%v", err.Error()),
//				Data:    nil,
//			}
//		}
//
//		return wxCilent.ResponseResult{
//			Code:    0,
//			Success: true,
//			Message: loginRes.GetBaseResponse().GetErrMsg().String(),
//			Data:    loginRes,
//		}
//	}
//
//	//30系列转向
//	if loginRes.GetBaseResponse().GetRet() == -301 {
//		return A16Login(Data, strings.Replace(string(loginRes.NetworkSectResp.BuiltinIPList.ShortConnectIPList[1].IP), string(byte(0x00)), "", -1), strings.Replace(string(loginRes.NetworkSectResp.BuiltinIPList.ShortConnectIPList[1].Domain), string(byte(0x00)), "", -1))
//	}
//
//	return wxCilent.ResponseResult{
//		Code:    int64(loginRes.GetBaseResponse().GetRet()),
//		Success: false,
//		Message: *loginRes.GetBaseResponse().GetErrMsg().String_,
//		Data:    loginRes,
//	}
//}
