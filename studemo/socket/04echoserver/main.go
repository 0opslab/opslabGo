package main

import (
	"bufio"
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
			reader := bufio.NewReader(conn)

			for {
				message, _ := reader.ReadString('\n')
				cmd := strings.TrimSpace(message)
				if cmd == "" {
					break
				}
				fmt.Println(conn.RemoteAddr().String() + " => " + cmd)
				if strings.EqualFold(cmd, "quit") {
					conn.Close()
				} else {
					echo := time.Now().String() + cmd + "\n"
					conn.Write([]byte(echo))
				}
			}

		}(conn)
	}
}
