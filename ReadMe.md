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

## Go语言编码规范
Go语言是严格区分大小写
Go应用程序的执行入口是main函数

####  Go的函数、变量、常量、自定义类型、包的命名方式遵循以下规则：
```
1）首字符可以是任意的Unicode字符或者下划线
2）剩余字符可以是Unicode字符、下划线、数字
3) 字符长度不限
4) golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，
    如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用可以简单的理解成，
    首字母大写是公有的，首字母小写是私有的
5) 结构体中属性名的大写
```

#### Go只有25个关键字
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var


## Go语言基础
### 变量声明初始化
变量声明以关键字 var 开头，后置变量类型，行尾无须分号，GO语言的变量声明格式为:
var 变量名  变量类型
```go
// 声明一个整形类型的变量，可以保存正数数值
var a int
// 声明一个字符串类型的变量
var b string
// 声明一个 32 位浮点切片类型的变量
var c []float32
// 声明一个返回值为布尔类型的函数变量
var d func() bool
//声明一个结构体类型的变量
var e struct{
	    x int
	}
//批量格式
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```
*变量初始化*的标准格式
var 变量名 类型 = 表达式
```go
//定义一个int类型值100的变量
var hp int = 100
//在标准格式的基础上，将 int 省略后，编译器会尝试根据等号右边的表达式推导 hp 变量的类型
var hp = 100

//简介类型的变量定义并初始化
hp := 100
//函数有两个返回值，一个是连接对象，一个是 err 对象
conn, err := net.Dial("tcp","127.0.0.1:8080")

// 多重赋值 a、b的值交换
b, a = a, b


```

在使用多重赋值时，如果不需要在左值中接收变量，可以使用匿名变量（anonymous variable）。
匿名变量的表现是一个下画线_，使用匿名变量时，只需要在变量声明的地方使用下画线替换即可
```go
func GetData() (int, int) {
    return 100, 200
}
a, _ := GetData()
_, b := GetData()
```

### 常量
Go 支持字符、字符串、布尔和数值 常量 ，const 用于声明一个常量。
```go
const s string = "constant"
```
### 流程控制
 Go 程序都是从 main() 函数开始执行，然后按顺序执行该函数体中的代码
 * if-else 结构
 * switch 结构
 * select 结构
 * for (range) 结构
下面是一些例子
```go
func main() {
    var first int = 10
    var cond int

    if first <= 0 {
        fmt.Printf("first is less than or equal to 0\n")
    } else if first > 0 && first < 5 {

        fmt.Printf("first is between 0 and 5\n")
    } else {

        fmt.Printf("first is 5 or greater\n")
    }
   
   var num1 int = 100
   
   switch num1 {
       case 98, 99:
           fmt.Println("It's equal to 98")
       case 100: 
           fmt.Println("It's equal to 100")
       default:
           fmt.Println("It's not equal to 98 or 100")
       }
   
   //一个 break 的作用范围为该语句出现后的最内部的结构，
   // 它可以被用于任何形式的 for 循环（计数器、条件判断等）。
   // 但在 switch 或 select 语句中，
   // break 语句的作用结果是跳过整个代码块，执行后续的代码。
   
   //关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，
   // 但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件
   for i:=0; i<5; i++ {
       for j:=0; j<10; j++ {
           println(j)
       }
   }
   
   
   //for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，
   // 即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）。
   LABEL1:
       for i := 0; i <= 5; i++ {
           for j := 0; j <= 5; j++ {
               if j == 4 {
                   continue LABEL1
               }
               fmt.Printf("i is: %d, and j is: %d\n", i, j)
           }
       }
}
```
### 指针
在GO语言中指针Pointer被拆分为俩个核心概念
* 类型指针，允许对这个指针类型的数据进行修改，传递数据使用指针，而无须数据拷贝。类型指针不能进行偏移和运算
* 切片，由指向初始元素的原始指针、元素数量和容量组成。

每个变量在运行时都拥有有一个地址，这个地址代表变量在内存中的位置，Go语言中使用&作符放在变量签名对变量进行“取地址”操作。
```go
// 其中v代表被取地址的变量，被取地址的v使用ptr变量进行接收，ptr的类型为*T ，称作T的执行类型。*代表指针类型
ptr := &v



var cat int = 1
var str string = "banana"
fmt.Printf("%p %p", &cat, &str)
//指针值带有0x的十六进制前缀
输出:0xc042052088 0xc0420461b0
```

### 数组
数组的写法如下:
```go
var 数组变量名 [元素数量]T
func main() {
    var arr1 [5]int

    for i:=0; i < len(arr1); i++ {
        arr1[i] = i * 2
    }

    for i:=0; i < len(arr1); i++ {
        fmt.Printf("Array at index %d is %d\n", i, arr1[i])
    }
}
```

### 切片slice
切片是Go中一个关键的数据类型，是一个比数组更加强大的序列号接口。类似于Python和Java中的list，在实际开发中使用的相当的多。
```go
//声明亲切的格式（不需要说明长度）
var name []type
//利用make函数创建一个切片
var slice1 []type = make([]type,len)
//简洁的写法
sclie2 := make([]type,len)
```
实例
```go
func main(){
	//声明一个切片
	s := make([]string,3)
	fmt.Println("emp:",s)
	
	//赋值
	s[0] ="a"
	s[1] ="b"
	s[2] ="c"
	
	//访问
	fmt.Println("set:",s)
	fmt.Println("get:",s[1])
	fmt.Println("len:",len(s))
	
	//动态追加元素
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("set:",s)
	
	
	//切片复制
	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("copy",c)
	
	//切片的切片语法
	//Slice 支持通过 slice[low:high] 语法进行“切片”操作。例如，这里得到一个包含元素 s[2], s[3],s[4] 的 slice。
    //这个 slice 从 s[0] 到（但是包含）s[5]
    
    l := s[2:5]
    fmt.Println("newSclie:",l)
	
}
```
将切片传递给函数
```go
func sum(a []int) int{
	s := 0
	for i :=0;i<len(a);i++{
		s += a[i]
	}
	return s
}
func main() {
    var arr = [5]int{0, 1, 2, 3, 4}
    sum(arr[:])
}

```

### 关联数组map
处理切片结果为map在日常开发中也经常会用到，尤其适用有关联关系的数据的时候。
map是引用类型，可以适用如下的声明方式
```go
var map1 map[keytype]valuetype
//简介的写法
map2 := make(map[string]int)
```
实例
```go
func main(){
	//声明一个map类型的变量
	m := make(map[string]int)
	
	
	//向变量内添加值
	m["k1"] = 1
	m["k2"] = 2
	
	//获取值
	fmt.Println(m["k1"])
	
	//删除值delete函数为内建函数
	delete(m["k2"])
	
	//另一种初始化方法
	n := map[string]int{"ss":1,"ba":2}
	
}
```

### Range遍历
在Go中有一种使用range可以迭代各种各样的数据结构。
```go
func main(){
	nums := []int[2,3,4]
	for k,v := range nums{
		fmt.Println("arr ",k," =>",v)
	}
	
	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
}
```

### 函数
函数是基本的代码块。Go是编译型的语言，所有函数编写的顺序是无关紧要的，一般把函数放在最前面，
函数的定义格式如下：
```go
//除了main()、init()函数外，其它所有类型的函数都可以有参数与返回值。
// 函数参数、返回值以及它们的类型被统称为函数签名

//函数名：由字母、数字、下画线组成。其中，函数名的第一个字母不能为数字。在同一个包内，函数名称不能重名。
//参数列表：一个参数由参数变量和参数类型组成
//返回参数列表：可以是返回值类型列表，也可以是类似参数列表中变量名和类型名的组合。
//      函数在声明有返回值时，必须在函数体中使用 return 语句提供返回值列表。
//函数调用
//      返回值变量列表 = 函数名(参数列表)
func 函数名(参数列表)(返回参数列表){
    函数体
}
```
在Go里面有三种类型的函数
* 普通的带有名字的函数
* 匿名函数或者lambda函数
* 方法（Methods）

实例
```go
/**
 * 定义一个函数
 */
func plus(a int,b int){
	return a + b
}

/**
 * 定义一个函数,返回多个值
 */
func vals()(int,int){
	return 0,1
}

/**
 * 定义一个函数,接收可变参数
 */
func min(a ...int) int{
	if len(a)==0 {
            return 0
    }
    min := a[0]
    for _, v := range a {
        if v < min {
            min = v
        }
    }
    return min
}


func main(){
	res := plus(1,2)
	fmt.Println("1+2=",res)
	
	//调用返回多个值的函数
	a,b := vals()
	
	//可变参数函数调用
	minvalue := min(1,2,0,3,4)
	fmt.Prinlnt(minvalue)
	
}
```
##### 匿名函数
匿名函数的定义格式如下:
```go
func (参数列表)(返回参数列表){
	函数体
}
//命名函数在定义时就可以调用
func(data string){
	fmt.Println("hello ",data)
}("Test")

//将匿名函数复制给变量
f := func(data string){
	fmt.Println("hello ",data)
}
//使用变量调用
f("Test")


//使用匿名函数用作回调函数
func visit(list []int,f func(int)){
	for _,v  := list{
		f(v)
	}
}

//使用匿名函数打印切片内容
visit([]int{1,2,3,4,5},func(v int){
	fmt.Println(v)
})

```

##### init函数
每一个源文件都可以有一个init函数，该函数会在main函数前被go的调度器钓鱼用。
该函数一般可以用来完成一些初始化操作。

##### 闭包
闭包就是一个函数和与其相关的环境(变量等)一同组合成有一个整体存在
闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含
在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提供绑定的计算环
境（作用域）。闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示
数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，就是说这些函数可以存储到
变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
```go
// 定义一个函数makeSuffix返回一个类型为func(string) string的函数
func makeSuffix(suffix string) func(string) string {
    return func(name string) string {
        if strings.HasSuffix(name, suffix) == false {
            return name + suffix
        }
        return name
    }
}

func main() {
    //判断字符串 以bmp结尾
    f1 := makeSuffix(".bmp")
    fmt.Println(f1("test"))
    fmt.Println(f1("pic"))
    f2 := makeSuffix(".jpg")
    fmt.Println(f2("test"))
    fmt.Println(f2("pic"))
}
```


##### defer延迟执行
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个
语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不
是单纯地返回某个值）。关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally
语句块，它一般用于释放某些已分配的资源
```go
func main() {
	prose_list = make([]string, 0)
	f, err := os.Open("c:/data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}
```
### defer panic recover
Go语言追求简洁优雅，所以Go语言不支持传统的try...catch..finally这种处理。Go中引入的处理方式为defer、panic、recover
使用panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
```go
func test(){
	defer func(){
		//使用defer+recover来捕获和处理异常
		err := recover()
		if err != nil{
			fmt.Println("err=",err)
		}
	}()
	
	num1 := 10
	num2 := 0
	res := num1 /num2
	fmt.Println("res=",res)
}
```
### 结构体
结构体是变量的集合，在数据组织的时候起着相当大的作用。
例如下面定义了一个结构体
```go
type Person struct{
	name string
	age  int
}
```
结构体的使用和常规的变量使用基本一致
```go
func main(){
	fmt.Println(Person{"Bob",20})
	fmt.Println(Person{name:"Bob",age:20})
	fmt.Println(Person{name:"Bob"})
	
	//指针类型
	ps := &Person{name:"Test",age：20}
	fmt.Println(ps.name)
	fmt.Println(ps.age)
	
	sp.name="Test1"
	sp.age=22
	fmt.Println(ps.name)
	fmt.Println(ps.age)
	
	//使用new方法创建执行结构体的指针
	t := new(Person)
	t.name ="test2"
	t.age = 20
	fmt.Println(t.name)
	fmt.Println(t.age)
	
	
	
}
```
为结构体添加方法
在Go语言中方法(Method)是一种做用于特定类型变量的函数。这种特定类型变量叫做接收器（Receiver）
如果将特定类型理解为结构体或类时，接收器的概念就类似于其他于语言中的this或self发方法。
```go
//接收器变量
//接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
//方法名、参数列表、返回参数：格式与函数定义一致。
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
```
实例
```go
//定义一个结构体
type rect struct {
    width, height int
}
//为结构体定义方法
func (r *rect) area() int {
    return r.width * r.height
}
//为结构体定义方法
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
```

### 接口
Go语言中虽然没有传统面向对象语言中类，集成的概念，不过提供了接口的支持，可以使用
接口来使用一些面向对象的特效。在Go语言中接口有如下特定
* 可以包含0个或多个方法的签名
* 只定义方法的签名，不包含实现
* 实现接口不需要显式的声明，只需实现相应方法即可
```go

//定义接口
type Shaper interface {
    Area() float32
    Circumference() float32
}

//定义结构体并且实现接口
type Rect struct {
    Width float32
    Height float32
}
func (r Rect) Area() int {
    return r.Width * r.Height
}

func (r Rect) Circumference() int {
    return 2 * (r.Width + r.Height)
}

//定义结构体并且实现接口
type Circle struct {
    Radius float32
}
func (c Circle) Area() int {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() int {
    return math.Pi * 2 * c.Radius
}

func showInfo(s Shaper) {
    fmt.Printf("Area: %f, Circumference: %f", s.Area(), s.Circumference())
}

//接口的调用
func main() {
    r := Rect{10, 20}
    showInfo(r)

    c := Circle{5}
    showInfo(r)    
}

//另外可以使用类型断言，来判断某一时刻接口是否是某个具体类型
v, ok := s.(Rect)   // s 是一个接口类型

//switch 不仅可以像其他语言一样实现数值、字符串的判断，还有一种特殊的用途——判断一个接口内保存或实现的类型。
func printType(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println(v, "is int")
    case string:
        fmt.Println(v, "is string")
    case bool:
        fmt.Println(v, "is bool")
    }
}
func main() {
    printType(1024)
    printType("pig")
    printType(true)
}
```

### 包
Go程序都只有一个文件，文件里包含一个main函数和几个其他的函数。包用于组织Go源代码，提供了更好的可重用性与可读写。


### Go的并发
#### goroutine
goroutine 的概念类似于线程，但 goroutine 由 Go 程序运行时的调度和管理。Go 程序会智能地将 goroutine 中的任务合理地分配给每个 CPU。
Go程序中可以适用go关键字为一个函数创建一个gorounite。一个函数可以被创建多个 goroutine，一个 goroutine 必定对应一个函数。

```go
func running() {
    var times int
    // 构建一个无限循环
    for {
        times++
        fmt.Println("tick", times)
        // 延时1秒
        time.Sleep(time.Second)
    }
}
func main() {
    // 并发执行程序
    // 适用go将一个普通函数转换为goroutine
    go running()
    // 接受命令行输入, 不做任何事情
    var input string
    fmt.Scanln(&input)
    
    //适用go将一个匿名函数转换为goroutine
     go func() {
            var times int
            for {
                times++
                fmt.Println("tick", times)
                time.Sleep(time.Second)
            }
     }()
}
```


### GO 的标准库
Go语言标准库包名	功  能
bufio	带缓冲的 I/O 操作
bytes	实现字节操作
container	封装堆、列表和环形列表等容器
crypto	加密算法
database	数据库驱动和接口
debug	各种调试文件格式访问及调试功能
encoding	常见算法如 JSON、XML、Base64 等
flag	命令行解析
fmt	格式化操作
go	Go 语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改
html	HTML 转义及模板系统
image	常见图形格式的访问及生成
io	实现 I/O 原始访问接口及访问封装
math	数学库
net	网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等
os	操作系统平台不依赖平台操作封装
path	兼容各操作系统的路径操作实用函数
plugin	Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载
reflect	语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值
regexp	正则表达式封装
runtime	运行时接口
sort	排序接口
strings	字符串转换、解析及实用函数
time	时间接口
text	文本模板及 Token 词法器


## 代码开发规范
说句真心话很多时候都觉得Goland就是个小孩的玩具，很多时候写着写着就爆发了。因此需要一个规范，来减少这种郁闷


### 代码注释
    //@Document 文档注释
    //     函数注释

## 在线教程
[Golang标准库文档](https://studygolang.com/pkgdoc)

[go入门教程](http://c.biancheng.net/golang/)

[go入门指南](https://books.studygolang.com/the-way-to-go_ZH_CN/)

[学习go的标准库](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/)

[go by example](https://books.studygolang.com/gobyexample/)


## 框架推荐
* [日志框架 logrus](https://github.com/sirupsen/logrus)

































