> # Quick Start

> ## Step 1：Install git
> ```
> git config --global --list
> git config --global user.name "snoredude"
> git config --global user.email "caffeine.dude.go@gmail.com"
> git config --global --unset http.cookiefile
> https://www.cnblogs.com/fireporsche/p/9359130.html
> https://www.cnblogs.com/dengxiaoning/p/13336783.html
> https://www.ipaddress.com/
> 140.82.113.3
> 199.232.69.194
> 140.82.113.3 github.com
> 199.232.69.194 github.global.ssl.fastly.net
> ipconfig /flushdns
> Unknown SSL protocol error in connection to github.com:443 
> 使用git从远程下载时，出现Unknown SSL protocol error in connection to xxx:443 错误。
> 很有可能是被墙在了外面，这里针对墙在外面的情况。
> git config --global http.proxy 127.0.0.1:1080
> git config --global http.sslVerify false
> https://blog.csdn.net/e_wsq/article/details/114262711
> ```

> ## Step 2：Install go

GOBIN
%GOPATH%\bin

GOPATH
E:\gopath

GOROOT
E:\goroot

打开cmd或者powershell，输入go version和go env，看到相关信息即配置成功

> ## Step 3：Install vscode
Tools environment: GOPATH=C:\Users\Administrator\go
Installing 10 tools at C:\Users\Administrator\go\bin in module mode.
 
在配置好环境已经，安装VSCode，之后安装Go插件，之后重启vscode。
接下来需要的就是在文件-首选项-设置中设置GoPath和GoRoot，这样才算是基本完成。

{
    "security.workspace.trust.untrustedFiles": "open",
    "go.useLanguageServer": false,
    "go.gopath": "E:\\gopath",
    "go.goroot": "E:\\goroot"
}

这里再说一点，你可以在VSCode安装一个Code Runner插件，直接右击就可以运行，当然你也可以在终端进入到hello.go文件目录，运行go install 或者 go build，之后在bin目录中找到hello.exe文件运行。结果是一样的。

https://blog.csdn.net/mnmiaoyi/article/details/98847144
https://blog.csdn.net/mnmiaoyi/article/details/98847144 镜像 网站

https://easydoc.net/p/57789218/u2Oy3U4O
https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
https://github.com/wechatpay-apiv3/wechatpay-go
https://blog.csdn.net/weixin_43314519/article/details/120556710

> ## Step 4：Start project
> ```
> git clone https://github.com/snoredude/doc.git
> cd /project/doc
> go version
> go env
> go env -w GO111MODULE=on
> go env -w GOPROXY=https://goproxy.cn,direct
> go mod init doc
> ```
> - go 版本 1.11 以上才能使用 go modules  
> - GO111MODULE=on 开启 go modules  
> - go env -w 会将配置信息写到环境变量 GOENV 中，可以用 `go env` 命令查看 go 的环境变量
> - 在 main.go 的同级目录下执行：`go mod init project_name` ，可以生成该项目的 go.mod 文件，  go.mod 会自动对项目进行依赖管理
> - 使用 go mod 管理项目时，依赖包存放在 `$GOPATH/pkg` 目录下，并且允许同一个依赖包有多个不同的版本存在
> - 多个项目可以共享 `$GOPATH/pkg` 目录下依赖包缓存
> - 在 windows 系统中，`$GOPATH/pkg` 目录的文件结构为：
>   1. mod
>   2. sumdb
>   3. windows_amd64

