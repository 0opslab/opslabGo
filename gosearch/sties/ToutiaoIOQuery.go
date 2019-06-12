package sties

import (
	"opslabGo/gosearch/entity"
	"time"
	"strconv"
	"opslabGo/gosearch/util"
	"github.com/PuerkitoBio/goquery"
)
/**
 * 爬去今日头条的搜索结果
 */
func TouTiaoQuery(key string, page int, timeout time.Duration, chan_result chan <- []entity.ResultInfo) {

	var arr  []entity.ResultInfo
	for i := 0; i < page; i++ {
		http_url := "https://toutiao.io/search?utf8=%E2%9C%93&q=" + key
		if (i > 0) {
			http_url += "&page=" + strconv.Itoa(i + 1)
		}

		doc := util.GoQueryRequestDocNoError(http_url, "https://toutiao.io", timeout)
		doc.Find(".post .content").Each(func(i int, s *goquery.Selection) {
			aTag := s.Find(".title a");
			title := aTag.Text();
			url, _ := aTag.Attr("href")
			desc := s.Find(".summary a").Text()

			result := entity.ResultInfo{
				From:"https://toutiao.io", Title:title, Info:"", Href:"https://toutiao.io/" + url,
				Desc:desc}
			arr = append(arr, result)

		})

	}
	chan_result <- arr
}

