package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
)

func getIPv6AddressFromApi() (string, error) {
	// 发送 GET 请求到 6.ipw.cn
	resp, err := http.Get("https://6.ipw.cn")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	// 读取 HTTP 响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将响应体转换成字符串并返回
	return string(bodyBytes), nil
}

func getIPv6AddressFromInterface() (string, error) {
	s, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range s {
		i := regexp.MustCompile(`(\w+:){7}\w+`).FindString(a.String())
		if strings.Count(i, ":") == 7 {
			return i, nil
		}
	}
	return "", errors.New("no ipv6 found")
}
