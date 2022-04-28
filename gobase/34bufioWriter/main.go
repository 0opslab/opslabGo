package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	buf_writer := bufio.NewWriter(os.Stdout)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := buf_writer.Write(d2)
	if err != nil{
		fmt.Println("write error :",err)
	}
	buf_writer.Flush()
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := buf_writer.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)
	buf_writer.Flush()

	n4, err := buf_writer.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	n5, err := buf_writer.WriteString("buffered 111 \n")
	fmt.Printf("wrote %d bytes\n", n5)

	buf_writer.Flush()
}
