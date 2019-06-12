package main

import (
	"fmt"
	"time"
)

//通过jobs来接收任务，通过result返回结果
func worker(id int, jobs <-chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("worker", id, "pricessing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启用工作池
	for w := 1; w < 10; w++ {
		go worker(w, jobs, results)
	}

	//派送任务
	for i := 0; i < 20; i++ {
		jobs <- i
	}
	//通过close表示任务派送完成
	close(jobs)

	//收集任务的返回值
	for i := 0; i < 20; i++ {
		<-results
	}

}
