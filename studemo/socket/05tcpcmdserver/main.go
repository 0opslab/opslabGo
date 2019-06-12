package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

/**
 * 利用go实现一个tcpcmdserver
 */
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}

		// start a new goroutine to handle the new connection
		go func(conn net.Conn) {
			//设置连接1分钟无数据关闭连接
			conn.SetReadDeadline(time.Now().Add(time.Minute * 1))
			defer conn.Close()
			conn.Write([]byte("$"))
			//循环处理请求
			for {
				buff := bufio.NewReader(conn)
				str, err := buff.ReadString('\n')
				if(err != nil){
					fmt.Println(err)
					conn.Write([]byte(err.Error()))
					return
				}
				//去掉两端空白字符
				cmd := strings.TrimSpace(str)
				//发送给客户端的数据
				fmt.Println("===>",cmd)
				rep := ""
				if cmd == "echo" {
					rep = "hello,client"
				} else if cmd == "time" {
					rep = time.Now().Format("2006-01-02 15:04:05")
				} else {
					res,err := Shellout("cmd",cmd)
					if(err != nil){
						rep += "ExecError:"+err.Error()
					}
					rep += string(res)
				}
				//发送数据
				rep += "$"
				conn.Write([]byte(rep))

			}
			_, err = conn.Write([]byte("server return :hello server"))
			return
		}(conn)
	}
}
func Shellout(shell ,command string) ( []byte ,error) {
	cmd := exec.Command(shell, "/c", command)
	return cmd.Output()
}