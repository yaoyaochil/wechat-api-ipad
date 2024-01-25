package Msg

type ShareCardParam struct {
	Wxid         string
	ToWxid       string
	Id     string
	NickName string
	Alias    string
}

type ShareLocationParam struct {
	Wxid		string
	ToWxid		string
	X 			float64
	Y			float64
	Scale		float64
	Title		string
	Poiname		string
}

type SendAppMsgParam struct {
	Wxid		string
	ToWxid		string
	Content			string
	Type 		int32
}
