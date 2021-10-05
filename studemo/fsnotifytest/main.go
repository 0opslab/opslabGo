package main

import (
	"opslabGo/studemo/fsnotifytest/Myfsnotify"
)
func main() {
	watch :=Myfsnotify.NewNotifyFile()
	watch.WatchDir("c:\\temp")

	//巧用如下方式实现阻塞
	select {}

	
}