package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg  sync.WaitGroup

	for i := 0;i< 100;i++ {
		wg.Add(1)

		go func(i int){
			defer wg.Done()
			fmt.Printf("goroutine end   %d \n",i)

		}(i)
	}

	//等待执行结果
	wg.Wait()
	fmt.Println("所有的goroutine执行完毕")
}