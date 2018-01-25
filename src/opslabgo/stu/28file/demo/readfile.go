package demo

import (
    "fmt"
	"os"
	"io"
	"bufio"
	//"encoding/json"
)


var prose_list []string

//func main() {
//	prose_list = make([]string,0)
//	f, err := os.Open("c:/data.json")
//    if err != nil {
//        panic(err)
//    }
//    defer f.Close()
//    rd := bufio.NewReader(f)
//    for {
//        line, err := rd.ReadString('\n')
//        if err != nil || io.EOF == err {
//            break
//		}
//		fmt.Println(line)
//	}
//
//	fmt.Println("===========")
//	fi, err := os.Open("c:/data.json")
//	if err != nil {
//        fmt.Printf("Error: %s\n", err)
//        return
//    }
//    defer fi.Close()
//
//    br := bufio.NewReader(fi)
//    for {
//        a, _, c := br.ReadLine()
//        if c == io.EOF {
//            break
//        }
//        fmt.Println(string(a))
//    }
//}