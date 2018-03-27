package sties

import (
	"testing"
	"opslabgo/gosearch/entity"
	"time"
	"fmt"
)

func TestCnblogsQuery(t *testing.T) {
	var key = "tomcat启动失败"

	chan_csdn := make(chan []entity.ResultInfo)
	go CnblogsQuery(key, 1, time.Second * 1, chan_csdn)

	select {


	case lst_csdn := <-chan_csdn:
		for _, v := range lst_csdn {
			fmt.Println(v)
		}
	}
}
