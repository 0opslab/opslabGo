package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)


func test(x int) {
	//获取锁
	cond.L.Lock() 
	fmt.Println("aaa: ", x)
	//等待通知  暂时阻塞
	cond.Wait()
	fmt.Println("bbb: ", x)
	time.Sleep(time.Second * 2)
	//释放锁
	cond.L.Unlock()
}


func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}

	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	// 下发一个通知给已经获取锁的goroutine
	cond.Signal()
	time.Sleep(time.Second * 1)
	// 3秒之后 下发一个通知给已经获取锁的goroutine
	cond.Signal()

	time.Sleep(time.Second * 1)
	//3秒之后 下发广播给所有等待的goroutine
	cond.Broadcast()
	time.Sleep(time.Second * 10)
	fmt.Println("finish all")

}
