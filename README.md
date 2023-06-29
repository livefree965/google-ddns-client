# Google Domain DDNS 转发服务客户端

本项目使用 Golang简单实现了客户端定时请求Google
Domain更新DNS，本服务主要结合[google-ddns-server](https://github.com/livefree965/google-ddns-server)
进行使用，目前仅支持ipv6地址更新和请求google-ddns-server服务进行更新，后续视需求进行进一步的开发。

## 项目背景

Google Domain 是 Google 提供的域名注册和管理服务，其中包括了动态域名解析（Dynamic DNS，简称 DDNS）功能，允许用户通过 API
更新域名的 IP 地址，实现动态域名的解析。然而，由于某些原因，Google Domain 在国内的访问受到限制，导致无法直接使用其 DDNS 接口。

为了解决这个问题，本项目使用 Golang 编写了一个简单的服务，请求由google-ddns-server代理的接口，实现了在国内也能够正常使用
Google Domain DDNS 服务的功能。

## 功能特点

- 将国内无法直接访问的 Google Domain DDNS 接口转发到可访问的地址，解决了无法使用 Google Domain DDNS 服务的问题。
- 使用 Golang 编写，具有高效性和跨平台特性，可以在多种操作系统上运行。
- 代码开源，您可以根据自己的需求进行修改和定制。

## 使用方法

1. 使用前需要先配置好google-ddns-server。

2. 在根目录下执行go build生成需要的二进制文件，也可以直接下载打包好的二进制包，修改config.ini的内容到你的配置并执行：

   ```bash
   ./google-ddns-client -c config.ini
   ```

3. 项目会通过接口每分钟查询一遍本地的ipv6地址，如果失败会请求ipw的接口尝试获取本地ipv6地址。

4. 项目另外简单打包支持了x86_64环境下的QNAP的应用安装包，将config.ini文件复制到安装目录下(一般是/share/CACHEDEV1_DATA/.qpkg/google-ddns-client)的/etc/config.ini后，在应用商店安装运行该脚本即可。

## 贡献

欢迎对本项目进行贡献！您可以通过以下方式参与项目：

- 提交问题和建议，帮助改进项目。
- 提交代码修复或新增功能，帮助增强项目。

## 许可

本项目使用 MIT 许可证，详细信息请参阅LICENSE文件。

## 帮助和支持

如有任何问题或需要帮助，请在 GitHub 项目页面提交问题，我们会尽快回复并解答您的疑问。
