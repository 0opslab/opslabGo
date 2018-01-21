package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	//定义个变量
	var ops uint64 = 0

	//启动50个协程来对机器数进行累积
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				//允许其他协程的执行
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second * 2)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
