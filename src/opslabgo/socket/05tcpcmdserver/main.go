package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

/**
 * 利用go实现一个tcpserver
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

			//循环处理请求
			for {
				data := make([]byte, 256)
				//从conn中读取数据
				n, err := conn.Read(data)
				//如果读取数据大小为0或出错则退出
				if n == 0 || err != nil {
					break
				}
				//去掉两端空白字符
				cmd := strings.TrimSpace(string(data[0:n]))
				//发送给客户端的数据
				rep := ""
				if cmd == "echo" {
					rep = "hello,client \n"
				} else if cmd == "time" {
					rep = time.Now().Format("2006-01-02 15:04:05")
				}
				//发送数据
				conn.Write([]byte(rep))

			}
			_, err = conn.Write([]byte("server return :hello server"))
			return
		}(conn)
	}
}
