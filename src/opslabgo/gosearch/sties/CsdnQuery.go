package sties

import (
"opslabgo/gosearch/entity"
"strconv"
"opslabgo/gosearch/util"
"github.com/PuerkitoBio/goquery"
	//	"log"
"time"
//	"github.com/google/logger"
)

/**
 * 爬去百度的搜索结果
 */
func CsdnQuery(key string, page int, timeout time.Duration, chan_result chan <- []entity.ResultInfo) {
	var arr  []entity.ResultInfo

	for i := 0; i < page; i++ {
		http_url := "https://so.csdn.net/so/search/s.do?q=" + key
		if (i > 0) {
			http_url += "&p=" + strconv.Itoa(i + 1) + "&t=blog&domain=&o=&s=&u=&l=&f=&rbg=0"
		}else{
			http_url += "&t=blog&o=&s=&l="
		}
		doc := util.GoQueryRequestDocNoError(http_url, "https://so.csdn.net", timeout)
		doc.Find(".search-list-con .search-list").Each(func(i int, s *goquery.Selection) {
			aTag := s.Find("dt").Find("a").First()
			href,_ := aTag.Attr("href")
			title := aTag.Text()
			desc := s.Find(".search-detail").Text()
			info := s.Find(".author-time").Text()
			result := entity.ResultInfo{
				From:"http://blog.csdn.net", Title:title, Info:info,
				Href:href, Desc:desc}
			arr = append(arr, result)
		})

	}
	chan_result <- arr
}

