package Mmtls

import (
	"net/http"
	"wechatwebapi/Cilent"
)

func GenNewHttpClient(Data *MmtlsClient) (httpclient *HttpClientModel) {

	mmtlsClient := &MmtlsClient{
		//不需要发送队列。
		ServerSeq: 1,
		ClientSeq: 1,
	}

	if Data != nil {
		mmtlsClient = Data
	}

	httpclientModel := &HttpClientModel{
		mmtlsClient: mmtlsClient,
		httpClient:  &http.Client{},
		curShortip:  Cilent.MMtls_host,
	}

	return httpclientModel
}
