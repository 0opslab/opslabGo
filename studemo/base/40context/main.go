package main

import (
	"context"
	"fmt"
	"time"
)


func main(){
	//创建一个context并存入一个指定的值
	ctx := context.WithValue(context.Background(),"trace_id",2222)
	//再想其中存入值
	ctx = context.WithValue(ctx,"session","sdlkfjkaslfsalfsafjalskfj")

	go func(ctx context.Context){
		ret,ok := ctx.Value("trace_id").(int)
		if !ok {
			ret = 21342423
		}

		fmt.Printf("ret:%d\n", ret)

		s , _ := ctx.Value("session").(string)
		fmt.Printf("session:%s\n", s)


	}(ctx)



	time.Sleep(1e10)
}
