package main

import (
	"fmt"
	"sync"
)

func main(){
	var once sync.Once
	onceBody := func(){
		fmt.Println("Run only once")
	}
	done := make(chan bool)

	for i:=0;i<10;i++{
		//虽然这段代码执行了10遍，但是onceBody只会执行一遍
		go func(){
			once.Do(onceBody)
			done <- true
		}()
		
	}

	for i:=0;i<10;i++{
		<-done
	}
}