package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func requestProxyGoogleDomain(url string, username string, password string, host string, ip string) string {
	completeUrl := fmt.Sprintf("%s?hostname=%s&ip=%s&username=%s&password=%s", url, host, ip, username, password)
	resp, err := http.Get(completeUrl)
	if err != nil {
		return err.Error()
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes)
}

func main() {
	consoleParams := parseNamedArgs()
	// 调用readConfigFile函数读取配置文件
	config, err := readConfigFile(consoleParams["c"])
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		return
	}
	for {
		// 获取IPv6地址
		ipv6Address, err := getIPv6AddressFromInterface()
		if err != nil {
			fmt.Println("Failed to get IPv6 address, Try use api:", err.Error())
			ipv6Address, err = getIPv6AddressFromApi()
		}
		if ipv6Address == "" {
			panic("can't find valid ipv6, exit")
		}
		fmt.Println("Local IPv6 Address:", ipv6Address)
		fmt.Println("update ddns result:", requestProxyGoogleDomain(config["url"], config["username"], config["password"], config["hostname"], ipv6Address))
		// 等待1分钟
		time.Sleep(time.Minute)
	}
}
