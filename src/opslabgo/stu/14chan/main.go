package main

import (
	"fmt"
)

//关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这
//个通道的接收方传达工作已经完成的信息。
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {

		for {
			//循环的从jobs 接收数据。在接收的这个特殊的二值形式的值中，如果 jobs 已经关闭了，
			//并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。
			//当我们完成所有的任务时，将使用这个特性通过 done 通道去进行通知
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done


	//通道的遍历
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
 