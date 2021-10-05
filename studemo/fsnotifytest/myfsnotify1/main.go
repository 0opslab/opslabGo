package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

//但节点监控
func main() {
  watcher, err := fsnotify.NewWatcher()
  if err != nil {
    log.Fatal("NewWatcher failed: ", err)
  }
  defer watcher.Close()

  done := make(chan bool)
  go func() {
    defer close(done)

    for {
      select {
      case event, ok := <-watcher.Events:
        if !ok {
          return
        }
				//event.Op 时间类型
				//CREATE
				//REMOVE
				//WRITE
				//RENAME
				//CHMOD
        log.Printf("%s %s\n", event.Name, event.Op)
      case err, ok := <-watcher.Errors:
        if !ok {
          return
        }
        log.Println("error:", err)
      }
    }
  }()

  err = watcher.Add("c:/temp")
  if err != nil {
    log.Fatal("Add failed:", err)
  }
  <-done
}