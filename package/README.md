```
$ mkdir $GOPATH/src/github.com/user/stringutil

//go build
$ go build github.com/user/stringutil

//go install
$ go install github.com/user/hello


bin/
    hello                 # 可执行命令
pkg/
    linux_amd64/          # 这里会反映出你的操作系统和架构
        github.com/user/
            stringutil.a  # 包对象
src/
    github.com/user/
            hello/
                hello.go      # 命令源码
    stringutil/
            reverse.go       # 包源码
```



