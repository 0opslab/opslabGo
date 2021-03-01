package main

import (
	"flag"
	"fmt"
)

//GO 提供了一个flag包，支持基本的命令行标志解析
func main() {

	// 声明一个默认值为foot的字符串标志word并带有一个简单的描述
	// 这里flag.String函数返回一个字符串指针

	wordPtr := flag.String("word", "foo", " a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}

/**
$ go build command-line-flags.go
word: opt
numb: 7
fork: true
svar: flag
tail: []
注意到，如果你省略一个标志，那么这个标志的值自动的设定为他的默认值。

$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
位置参数可以出现在任何标志后面。

$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]
注意，flag 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）。

$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]
使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。

$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
如果你提供一个没有使用 flag 包指定的标志，程序会输出一个错误信息，并再次显示帮助文本。

$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
 */