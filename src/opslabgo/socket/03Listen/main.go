package main

import (
	"net"
	"log"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	checkErr(err)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					log.Println(err)
					return
				}

				c.Write(buf[:n])
			}
		}(conn)

	}

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}