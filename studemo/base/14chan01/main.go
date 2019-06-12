package main

import "fmt"

func print(c chan int){
	for{
		//开始无限循环等待数据

		//从channel中读取数据并打印,当数据为0的时候退出循环
		data := <- c

		if data == 0 {
			break
		}
		fmt.Println(data)
	}

	//通知调用者打印完毕
	c <- 0
}

func main(){

	//创建channel
	ch := make(chan int)

	go print(ch)

	for i:= 1;i<= 10; i++{
		//通过通道将数据传递给print
		ch <- i
	}

	//通知print数据传完了
	ch <- 0

	//等等print执行完毕
	_,ok := <- ch
	fmt.Println(ok)
}
