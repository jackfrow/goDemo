package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 服务器地址
	serverURL := "https://example.com:443"

	// 创建默认的 Transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}

	// 创建自定义 Client
	client := &http.Client{Transport: tr}

	// 发起 HTTPS 请求
	resp, err := client.Get(serverURL)
	if err != nil {
		fmt.Println("Failed to perform HTTPS request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response from server:")
	fmt.Println(string(body))
}