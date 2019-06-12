package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ScanPort(ip string, portlist []int, result chan string) {
	//log.Println("port list:",portlist)

	result_port := ""
	for i := 0; i < len(portlist); i++ {
		_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(portlist[i]), time.Second*3)
		if err != nil {
			continue
		}
		result_port += strconv.Itoa(portlist[i]) + ","
	}
	result <- strings.Trim(result_port, ",")

}

func main() {
	result := make(chan string)
	limit := 1
	var ports = []int{22, 80, 135, 145, 8080}

	for i := 0; i < limit; i++ {
		// 一个ip一个携程
		go ScanPort("127.0.0.1", ports, result, )
	}
	println("127.0.0.1 open port:", <-result)


	//扫描全部端口
	result1 := make(chan string)
	pport := 1
	for i := 0; i < 5000; i++ {
		all_port := make([]int, 14)
		for j := 0; j < 14; j++ {
			pport++
			all_port[j] = pport
		}
		go ScanPort("127.0.0.1", all_port, result1)
	}

	res_port := []int{}
	for i := 0; i < 5000; i++ {
		strTemp := <-result1
		if strTemp != "" {
			tt := strings.Split(strTemp, ",")
			for _, value := range tt {
				port, _ := strconv.Atoi(value)
				res_port = append(res_port, port)
			}
		}
	}
	sort.Ints(res_port)
	ss := strings.Replace(strings.Trim(fmt.Sprint(res_port), "[]"), " ", ",", -1)
	println("127.0.0.1 open port:", ss)

	close(result1)

}
