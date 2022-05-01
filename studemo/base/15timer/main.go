package main

import (
	"fmt"
	"time"
)

func main() {
	//新建计时器，俩秒以后触发，
	//go触发计时器的方法比较特别，就是在计时器的channle中发送值

	timer1 := time.NewTimer(time.Second * 2)
	//此处等待channel中的信号，执行此段代码会阻塞俩秒
	<-timer1.C
	fmt.Println("timer 1 expried")

	//新建一个计时器3秒后触发
	timer2 := time.NewTimer(time.Second * 3)
	//新开一个协程来处理触发后的事件
	//因为是新起协程处理，因此不会阻塞
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()

	//停止
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

}
