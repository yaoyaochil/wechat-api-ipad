package Login

import (
	"fmt"
	wxCilent "wechatwebapi/Cilent"
	"wechatwebapi/comm"
)

func CacheInfo(Wxid string) wxCilent.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return wxCilent.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	return wxCilent.ResponseResult{
		Code:    1,
		Success: true,
		Message: "成功",
		Data:    D,
	}
}
