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

> ## Step 3：Install vscode

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

