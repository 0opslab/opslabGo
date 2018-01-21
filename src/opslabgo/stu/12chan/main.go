package main

import (
	"fmt"
	"time"
)

//通道是连接多个go协程的通道，可以从一个go协程将值发送到通道，然后在别的协程中接收
//默认发送和接收操作都是阻塞的，直到发送和接收放都准备完毕

//使用 channel <- 语法发送一个新的值到通道中
//使用 <- channel 语法从通道中读取一个值

//通道默认是无缓冲的。可以在创建的时候设定缓冲的个数
func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	channel := make(chan string, 2)
	channel <- "buffered"
	channel <- "channel"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	done := make(chan bool, 1)
	go worker(done)

	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//通过通道来告知调用者，函数执行完毕
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("worked")

	//通过通道来告知执行完毕
	done <- true
}

//定义该函数中允许发送数据，不允许接收数据
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//定义函数pings用来接收数据，pongs用来发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
