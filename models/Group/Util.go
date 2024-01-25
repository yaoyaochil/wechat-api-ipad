package Group

type AddChatRoomParam struct {
	Wxid    string
	ToWxids string
	QID     string
}

type GetChatRoomParam struct {
	Wxid string
	QID  string
}

type OperateChatRoomInfoParam struct {
	Wxid    string
	QID     string
	Content string
}

type MoveContractListParam struct {
	Wxid string
	QID  string
	Val  uint32
}

type OperateChatRoomAdminParam struct {
	Wxid    string
	QID     string
	ToWxids string
	Val     int32
}
