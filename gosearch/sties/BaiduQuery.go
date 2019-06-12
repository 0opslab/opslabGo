package sties

import (
	"opslabGo/gosearch/entity"
	"strconv"
	"opslabGo/gosearch/util"
	"github.com/PuerkitoBio/goquery"
	"time"
)

/**
 * 爬去百度的搜索结果
 */
func BaiduGuoQuery(key string, page int, timeout time.Duration, chan_result chan <- []entity.ResultInfo) {
	var arr  []entity.ResultInfo

	page_count := 10
	for i := 0; i < page; i++ {
		http_url := "https://www.baidu.com/s?ie=utf-8&wd=" + key
		if (i > 0) {
			http_url += "&pn=" + strconv.Itoa(i * page_count) + "&usm=" + strconv.Itoa(i + 1)
		}
		doc := util.GoQueryDoc(http_url, timeout)
		doc.Find(".result").Each(func(i int, s *goquery.Selection) {
			aTag := s.Find("h3[class=t]").Find("a").First()
			url,_ := aTag.Attr("href");
			title := aTag.Text()
			desc := s.Find("div[class=c-abstract]").Text()

			result := entity.ResultInfo{
								From:"https://www.baidu.com", Title:title, Info:"",
							Href:url, Desc:desc}
							arr = append(arr, result)
			// for m.baidu.com
			//data_log, ok := s.Attr("data-log")
			//if ok {
			//	m := make(map[string]interface{})
			//
			//	err := json.Unmarshal([]byte(strings.Replace(data_log, "'", "\"", -1)), &m)
			//	if err == nil {
			//		url := m["mu"]
			//		switch v := url.(type) {
			//		case string:
			//			//fmt.Println(v)
			//			title := strings.Replace(s.Find("h3").Text(), "\"", "'", -1)
			//			desc := s.Find(".c-abstract").Text()
			//			create := s.Find("span[class=c-gray]").Text()
			//
			//			result := entity.ResultInfo{
			//				From:"https://www.baidu.com", Title:title, Time:create,
			//				Href:v, Desc:strings.Replace(desc, "\"", "'", -1)}
			//			arr = append(arr, result)
			//		}
			//	}
			//
			//}
		})

	}
	chan_result <- arr
}
