package main

import (
	"os"
	"fmt"
	"net"
)


//IP地址是现在互联网的通信基础
//在GO的net包中定义了该类型和很多网络编程相关的函数
func main() {
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr,"Usage:%s ip-add\n",os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr != nil{
		fmt.Println("The address is",addr.String())
	}else {
		fmt.Println("Invalid address")
	}
	os.Exit(0)
}
