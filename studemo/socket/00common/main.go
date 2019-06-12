package main

import (
	"fmt"
	"net"
	"os"
)

//@Document演示net包中常用的函数使用

func main(){

	//将string类型的ip地址转为IP对象并根据ip获取子网掩码等

	ipStr := "192.168.30.10"

	ip := net.ParseIP(ipStr)

	if ip == nil{
		fmt.Println("无效的地址")
		return
	}

	defaultMask := ip.DefaultMask()
	fmt.Println( "DefaultMask:", defaultMask, defaultMask.String())

	ones, bits := defaultMask.Size()
	fmt.Println("ones: ",ones," bits: " , bits)

	network := ip.Mask(defaultMask)
	fmt.Println(os.Stdout, "network:", network.String())



	//通过域名获取Ip
	domain := "www.baidu.com"
	ipAddr,err := net.ResolveIPAddr("ip",domain)
	if err != nil{
		fmt.Printf("域名解析IP异常",err)
		return
	}
	fmt.Println(domain,"===>",ipAddr)

	//动态dns查询域名对应的所有ip地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Println( "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Println(n)
	}

	// 查看telnet服务器使用的端口
	port, err := net.LookupPort("tcp", "telnet")

	if err != nil {
		fmt.Println("未找到指定服务")
		return
	}

	fmt.Println( "telnet port: ", port)

	// 将一个host地址转换为TCPAddr。host=ip:port
	pTCPAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Err: ", err.Error())
		return
	}
	fmt.Printf( "www.baidu.com:80 IP: %s PORT: %d", pTCPAddr.IP.String(), pTCPAddr.Port)
}
