package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var rw sync.RWMutex


func read(n int,ch chan struct{}){
	rw.RLock()
    v := count
    fmt.Printf("进入读操作...值为：%d\n", v)
    rw.RUnlock()
    ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
    rw.Lock()
    v := rand.Intn(1000)
    count = v
    fmt.Printf("===>进入写操作 新值为：%d\n", v)
    rw.Unlock()
    ch <- struct{}{}
}


func main(){
	ch := make(chan struct{},10)
    for i := 0; i < 10; i++ {
        go read(i, ch)
    }
    for i := 0; i < 5; i++ {
        go write(i, ch)
    }

    for i := 0; i < 10; i++ {
        <-ch
    }

}