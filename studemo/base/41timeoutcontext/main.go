package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx,_ := context.WithTimeout(context.Background(),2 * time.Second)
	go testA(ctx)

	time.Sleep(5 * time.Second)
}

func testA(ctx context.Context){
	ch := make(chan string)
	select {
	case <-ctx.Done():
		fmt.Println("ctx.Done timeout")
		//ch <- "Time out"
		return
	case i := <-ch:
		fmt.Println(i)
	}
}


