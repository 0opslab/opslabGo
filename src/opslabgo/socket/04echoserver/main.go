package main

import (
	"bufio"
	"fmt"
	"net"
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

				if string(message) == "quit" {
					conn.Close()
				} else {
					echo := time.Now().String() + string(message) + "\n"
					conn.Write([]byte(echo))
				}
			}

		}(conn)
	}
}
