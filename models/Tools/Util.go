package Tools

type UploadParam struct {
	Wxid	string
	Base64	string
}

type DownloadVoiceParam struct {
	Wxid		string
	FromUserName	string
	NewMsgId	string
	Bufid		string
	Length		int
}

type DownloadVoiceData struct {
	Base64 []byte
	VoiceLength int32
}
