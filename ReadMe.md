
# 目录说明
## bin 
bin文件夹存放go install命名生成的可执行文件，可以把$GOPATH/bin路径加入到PATH环境变量里
就和我们上面配置的$GOROOT/bin一样，这样就可以直接在终端里使用我们go开发生成的程序了。

# pkg
pkg 文件是存放go编译生成的文件

# src
src是存放go源代码，不同的工程项目的代码以包名区分。


# 交叉编译
Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用，这里备忘一下。

## Mac 下编译 Linux 和 Windows 64位可执行程序
```bash
CGO_ENABLED=0 
GOOS=linux 
GOARCH=amd64 
go build main.go


CGO_ENABLED=0 
GOOS=windows 
GOARCH=amd64 
go build main.go
```
## Linux 下编译 Mac 和 Windows 64位可执行程序
```bash
CGO_ENABLED=0 
GOOS=darwin 
GOARCH=amd64 
go build main.go


CGO_ENABLED=0 
GOOS=windows 
GOARCH=amd64 
go build main.go
```
Windows 下编译 Mac 和 Linux 64位可执行程序
```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

## 安装库
话说go还在成长阶段,如果那go开发的任务比较重的话，还需要做很多工作，安装很多的第三方库，当然安装第三方库很简单。
下面是是安装一个生成UUID的库。
```bash
go get github.com/satori/go.uuid
```


### 现在问的
[go入门教程](http://c.biancheng.net/golang/)

[go入门指南](https://books.studygolang.com/the-way-to-go_ZH_CN/)

[学习go的标准库](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/)

[go by example](https://books.studygolang.com/gobyexample/)