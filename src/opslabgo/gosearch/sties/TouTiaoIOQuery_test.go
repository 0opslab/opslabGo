package sties

import (
	"testing"
	"opslabgo/gosearch/entity"
	"time"
	"fmt"
)

func TestTouTiaoQuery(t *testing.T) {
	var key = "tomcat启动失败"

	chan_test := make(chan []entity.ResultInfo)
	go CsdnQuery(key, 2, time.Second * 1, chan_test)

	select {


	case lst_csdn := <-chan_test:
		for _, v := range lst_csdn {
			fmt.Println(v)
		}
	}
}
