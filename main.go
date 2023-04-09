package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func getIPv6Address() (string, error) {
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

// 函数用于解析带有名字的命令行参数，并返回参数名和参数值的映射关系
func parseNamedArgs() map[string]string {
	// 获取命令行参数
	args := os.Args

	// 定义一个map用于存储参数名和参数值
	params := make(map[string]string)

	// 遍历命令行参数
	for i := 1; i < len(args); i++ {
		arg := args[i]

		// 检查参数是否以"-"开头
		if strings.HasPrefix(arg, "-") {
			// 提取参数名和参数值
			name := arg[1:]
			value := ""

			// 如果有下一个参数，并且下一个参数不是以"-"开头，则将其作为参数值
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				value = args[i+1]
				i++ // 跳过下一个参数，因为它已经作为参数值被处理了
			}

			// 存储参数名和参数值到map中
			params[name] = value
		}
	}

	return params
}

func requestProxyGoogleDomain(url string, username string, password string, host string, ip string) string {
	completeUrl := fmt.Sprintf("%s?hostname=%s&ip=%s&username=%s&password=%s", url, host, ip, username, password)
	resp, err := http.Get(completeUrl)
	if err != nil {
		return err.Error()
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes)
}

// readConfigFile 读取配置文件，并返回解析后的键值对的map
func readConfigFile(filename string) (map[string]string, error) {
	config := make(map[string]string)

	// 打开配置文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// 创建Scanner以逐行读取配置文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 跳过空行和注释（以 "#" 开头的行）
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析键值对
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		} else {
			return nil, fmt.Errorf("invalid config line: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	return config, nil
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
		// 获取主机名
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Failed to get hostname:", err)
			continue
		}
		// 获取IPv6地址
		ipv6Address := ""
		addrs, err := net.LookupIP(hostname)
		if err != nil {
			fmt.Println("Failed to get IPv6 address, Try use api:", err.Error())
			ipv6Address, _ = getIPv6Address()
		} else {
			for _, addr := range addrs {
				if ipv6 := addr.To16(); ipv6 != nil && ipv6.IsGlobalUnicast() && ipv6.To4() == nil {
					ipv6Address = ipv6.String()
					break
				}
			}
		}

		if ipv6Address != "" {
			fmt.Println("Local IPv6 Address:", ipv6Address)
			fmt.Println("update ddns result:", requestProxyGoogleDomain(config["url"], config["username"], config["password"], config["hostname"], ipv6Address))
		} else {
			fmt.Println("Failed to get IPv6 address.")
		}

		// 等待1分钟
		time.Sleep(time.Minute)
	}
}
