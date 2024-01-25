package Login

import (
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/comm"
)

func Get62Data(Wxid string) string {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return err.Error()
	}
	return wxCilent.Get62Data(D.Deviceid_str)
}
