package sites

import (
	"opslabgo/gosearch/entity"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func Baidu_search(key string, count int) ([]entity.ResultInfo, error) {
	result := make([]entity.ResultInfo,count)

	doc, err := goquery.NewDocument("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=0&rsv_idx=1&tn=baidu&wd=goquery")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div").Each(func(i int, contentSelection *goquery.Selection) {
		fmt.Println(contentSelection.Text());
	})


	result[0] = entity.ResultInfo{
		"2017/10",
		"http://bing.com",
		"测试"}

	return result, nil
}
