package main

import (
	"encoding/json"
	"flag"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"opslabGo/gosearch/entity"
	"opslabGo/gosearch/sties"
	"opslabGo/gosearch/util"
	"sort"
	"strings"
	"time"
)

const logPath = "/tmp/gosearch.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {

	flag.Parse()

	//lf, err := os.OpenFile(logPath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0660)
	//if err != nil {
	//	logger.Fatalf("Failed to open log file: %v", err)
	//}
	//defer lf.Close()
	//
	//defer logger.Init("LoggerExample", *verbose, true, lf).Close()


	//设置访问的路由
	http.HandleFunc("/s.php", Search)

	//设置监听的端口
	errs := http.ListenAndServe(":10000", nil)
	if errs != nil {
		logger.Info("start error :", errs)
	}
	logger.Info("[Info] service is started ... ")

}

func Search(w http.ResponseWriter, r *http.Request) {

	var rst_lst []entity.ResultInfo
	r.ParseForm()
	key := r.Form.Get("k")
	if (key != "") {
		logger.Info("Access", r.RequestURI)
		request_key := util.UrlEncode(key)
		chan_baidu := make(chan []entity.ResultInfo)
		go sties.BaiduGuoQuery(request_key, 2, time.Second * 1, chan_baidu)

		chan_toutiaoio := make(chan []entity.ResultInfo)
		go sties.TouTiaoQuery(request_key, 2, time.Second * 1, chan_toutiaoio)

		chan_csdn := make(chan []entity.ResultInfo)
		go sties.CsdnQuery(request_key, 2, time.Second * 1, chan_csdn)

		chan_cnblogs := make(chan []entity.ResultInfo)
		go sties.CnblogsQuery(request_key, 2, time.Second * 1, chan_cnblogs)

		for i := 0; i < 4; i++ {
			select {
			case lst_baidu := <-chan_baidu:
				for _, v := range lst_baidu {
					rst_lst = append(rst_lst, v)
				}
			case lst_toutiao := <-chan_toutiaoio:
				for _, v := range lst_toutiao {
					rst_lst = append(rst_lst, v)
				}

			case lst_csdn := <-chan_csdn:
				for _, v := range lst_csdn {
					rst_lst = append(rst_lst, v)
				}
			case lst_cnblog := <-chan_csdn:
				for _, v := range lst_cnblog {
					rst_lst = append(rst_lst, v)
				}
			case <-time.After(time.Second * 2):
				continue
			}

		}

	}

	var lst []entity.ResultInfo
	remove_duplicate := make(map[string]string)

	filer_sties := [...]string{
		"https://m.baidu.com"}

	for _, v := range rst_lst {
		if v.Href == "" || len(v.Href) <= 0 {
			continue
		}
		for _, filter := range filer_sties {
			if strings.HasPrefix(v.Href, filter) {
				continue
			}
		}
		_, ok := remove_duplicate[v.Href]
		if ! ok {
			remove_duplicate[v.Href] = v.Href
			entity.SetOrder(&v, key)
			lst = append(lst, v)
		}

	}

	sort.Sort(entity.ResultSlice(lst))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	b, _ := json.Marshal(lst)
	fmt.Fprintf(w, string(b))

}

