package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
