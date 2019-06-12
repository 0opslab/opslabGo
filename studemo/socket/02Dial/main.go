package main

import (
	"net"
	"io/ioutil"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":2000")
	checkErr(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)

	result, err := ioutil.ReadAll(conn)
	checkErr(err)

	fmt.Println(string(result))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}