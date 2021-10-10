## 简要说明 
这个库是学习和测试用的，因此很缭乱。不要企图再里面能找到有用的完整的代码，都是些小小的玩具。仔细阅读或许能发现一些或许有用的东西吧。祝你好运。

## 安装golang.org/x/库
虽说GO1.12开始支持go module了，但是由于某些原因比如操蛋的网络问题，搞的拉个库和拉屎一样，因此有些库不得不手动安装。
```bash
$mkdir -p $GOPATH/src/golang.org/x/
$cd $GOPATH/src/golang.org/x/
$git clone https://github.com/golang/net.git net
$go install net

$#download sync text tools crypto
git clone https://github.com/golang/sync.git sync
go install sync

git clone https://github.com/golang/sync.git text
go install text

git clone https://github.com/golang/sync.git crypto
go install crypto

git clone https://github.com/golang/sync.git tools
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install golang.org/x/tools/cmd/fiximports
go install golang.org/x/tools/cmd/gopls
go install golang.org/x/tools/cmd/godex

```

## 安装库
话说go还在成长阶段,如果那go开发的任务比较重的话，还需要做很多工作，安装很多的第三方库，当然安装第三方库很简单。
下面是是安装一个生成UUID的库。
```bash
go get github.com/satori/go.uuid
```

## 常用命令
```bash
build: 编译包和依赖
clean: 移除对象文件
doc: 显示包或者符号的文档
env: 打印go的环境信息
bug: 启动错误报告
fix: 运行go tool fix
fmt: 运行gofmt进行格式化
generate: 从processing source生成go文件
get: 下载并安装包和依赖
install: 编译并安装包和依赖
list: 列出包
run: 编译并运行go程序
test: 运行测试
tool: 运行go提供的工具
version: 显示go的版本
vet: 运行go tool vet
```


## go代码添加版本
通过使用版本标签来发布1.0.0版本：
git tag v1.0.0
git push --tags
此时将会在Github的仓库上创建名为v1.0.0的标签。推荐的做法是创建新的代码分支，这样可以直接在分支上修改v1.0.0的问题，而不影响主分支的开发进度。

git checkout -b v1
git push -u origin v1



























