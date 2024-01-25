package Mmtls

import (
	"bytes"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"wechatwebapi/models"
)

func (httpclient *HttpClientModel) POST(ip string, cgiurl string, data []byte, host string, P models.ProxyInfo) ([]byte, error) {
	var dialer proxy.Dialer
	var iphost string
	var err error
	iphost = "http://"
	iphost += ip
	iphost += cgiurl
	body := bytes.NewReader(data)

	var c *http.Client
	//设定代理
	if P.ProxyIp != "" {

		var ProxyUser *proxy.Auth

		if P.ProxyUser != "" && P.ProxyPassword != "" {
			ProxyUser = &proxy.Auth{
				User:     P.ProxyUser,
				Password: P.ProxyPassword,
			}
		} else {
			ProxyUser = nil
		}

		dialer, err = proxy.SOCKS5("tcp", P.ProxyIp, ProxyUser, proxy.Direct)

		if err != nil {
			return []byte{}, err
		}

		c = &http.Client{
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
			Timeout: time.Second * 30,
		}
	} else {
		c = &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * 3)) //设置发送接受数据超时
					return conn, nil
				},
				ResponseHeaderTimeout: time.Second * 3,
			},
		}
	}

	request, err := http.NewRequest("POST", iphost, body)
	if err != nil {
		return []byte(""), err
	}
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "close")
	request.Header.Set("Content-type", "application/octet-stream")
	request.Header.Set("User-Agent", "MicroMessenger Client")
	request.Close = true
	var resp *http.Response
	resp, err = c.Do(request)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}
