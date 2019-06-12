package main

import (
	"fmt"
	"net"
)

//@Document 查看系统的网络接口信息
func main(){
	inters,err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网络接口信息错误:",err)
		return
	}
	for _,v := range(inters){
		fmt.Printf("%+v\n",v)
	}


	//返回网络接口地址信息
	addrs,err := net.InterfaceAddrs()
	if err != nil{
		fmt.Println("获取网络接口信息错误:",err)
		return
	}
	for _,v := range(addrs){
		fmt.Printf("%+v\n",v)
	}
}
