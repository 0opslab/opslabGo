package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	//mutex将同步对state的方法
	var mutex = &sync.Mutex{}

	var ops int64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 20; i++ {
		go func() {
			key := rand.Intn(5)
			value := rand.Intn(100)
			mutex.Lock()
			state[key] = value
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second * 10)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	//对 state 使用一个最终的锁，显示它是如何结束的。

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
