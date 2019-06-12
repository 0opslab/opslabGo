package sties

import (
	"testing"
	"opslabGo/gosearch/entity"
	"time"
	"fmt"
)

func TestBaiduGuoQuery(t *testing.T) {
	var key = "tomcat启动失败"

	chan_test := make(chan []entity.ResultInfo)
	go BaiduGuoQuery(key, 1, time.Second * 1, chan_test)

	select {


	case lst_csdn := <-chan_test:
		for _, v := range lst_csdn {
			fmt.Println(v)
		}
	}
}
