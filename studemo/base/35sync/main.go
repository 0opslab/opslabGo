package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	a := 0

	//通过锁控制并发资源的访问
	var lock  sync.Mutex
	for i := 0;i<100;i++{
		go func(i int){
			//获取锁
			lock.Lock()
			defer lock.Unlock()

			
			a += i
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(a)
}