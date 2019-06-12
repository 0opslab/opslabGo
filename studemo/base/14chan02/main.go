package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

//产生数据
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump2(ch chan string) {
	for i := 0; ; i++ {
		ch <- "Strings" + strconv.Itoa(i)
	}
}

//消费数据
func suck(ch1 chan int, ch2 chan string) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received on channel 1: ", v)
		case v := <-ch2:
			fmt.Println("Received on channel 2: ", v)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 2")

		//可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select
		//default:
		//	fmt.Println("no message received")
		}
	}
}
