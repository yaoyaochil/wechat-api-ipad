package Mmtls

import "wechatwebapi/Fun"

//分包
func Separate(Data []byte) [][]byte {
	var NewData [][]byte
	for {
		if len(Data) > 0 {
			Len  := Data[3:5]
			NewData = append(NewData,Data[6:int64(Fun.Hex2int(&Len))])
			Data = Data[5 + int64(Fun.Hex2int(&Len)):]
		}else{
			break
		}
	}
	return NewData
}