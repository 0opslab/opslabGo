package main

import (
	"time"
	"fmt"
)

func main() {
	//go的通道选取可以同时等待多个通道操作。
	//go协程和统一以及选取的结合是go的一个强大特性

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout ")
		}
	}

	//常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default
	// 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
	messages := make(chan string)
	signals := make(chan bool)

	//这里是一个非阻塞接收的例子。如果在 messages 中存在，然后 select
	//将这个值带入 <-messages case中。如果不是，就直接到 default 分支中。

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

