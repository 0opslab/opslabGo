package main

import (
	"os"
	"fmt"
)

func main() {
	//os.Args提供原始命令行参数访问功能
	argsWithProg := os.Args
	fmt.Println(argsWithProg)


	argsWithOutProg := os.Args[1:]
	fmt.Println(argsWithOutProg)
	//可以使用标准的索引为止方式来取得单个参数的值
	arg := os.Args[3]
	fmt.Println(arg)




}
