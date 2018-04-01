package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	//从命令行读取内容并且打印
	scanner  := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input_str := scanner.Text()
		if(input_str == "quit"){
			os.Exit(0)
		}else{
			fmt.Println(input_str)
		}

	}

	if err := scanner.Err();err != nil{
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
