package bts

import (
	"encoding/json"
	"wechatwebapi/Cilent/mm"
)

func SearchContactResponse(Data interface{}) mm.SearchContactResponse {
	var Buff mm.SearchContactResponse
	result, err := json.Marshal(&Data)
	if err != nil {
		return mm.SearchContactResponse{}
	}
	_ = json.Unmarshal(result, &Buff)

	return Buff
}

func GetContactResponse(Data interface{}) mm.GetContactResponse {
	var Buff mm.GetContactResponse
	result, err := json.Marshal(&Data)
	if err != nil {
		return mm.GetContactResponse{}
	}
	_ = json.Unmarshal(result, &Buff)
	return Buff
}

func GetModContact(Data interface{}) mm.ModContact {
	var Buff mm.ModContact
	result, err := json.Marshal(&Data)
	if err != nil {
		return mm.ModContact{}
	}
	_ = json.Unmarshal(result, &Buff)
	return Buff
}
