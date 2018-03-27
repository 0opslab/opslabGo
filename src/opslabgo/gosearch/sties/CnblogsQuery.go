package sties




import (
"opslabgo/gosearch/entity"
"strconv"
"opslabgo/gosearch/util"
"github.com/PuerkitoBio/goquery"
"time"
//	"github.com/google/logger"
"strings"
	"regexp"
)

/**
 * 爬去百度的搜索结果
 */
func CnblogsQuery(key string, page int, timeout time.Duration, chan_result chan <- []entity.ResultInfo) {
	var arr  []entity.ResultInfo

	for i := 0; i < page; i++ {
		http_url := "http://zzk.cnblogs.com/s?t=b&w=" + key
		if (i > 0) {
			http_url += "http://zzk.cnblogs.com/s/blogpost?Keywords="+key+"&pageindex="+ strconv.Itoa(i + 1)
		}
		cookies := map[string]string{
			"ga": "GA1.2.2090750614.1514033118",
			"UM_distinctid":"16083684c38644-0f36ab5bb448c9-163c6657-1fa400-16083684c3951d",
			"_gid":"GA1.2.27017475.1521816028",
			"__utmc":"59123430",
			"__utmz":"59123430.1521948789.1.1.utmcsr=cnblogs.com|utmccn=(referral)|utmcmd=referral|utmcct=/",
			"__utma":"59123430.2090750614.1514033118.1521953424.1521956819.3"}



		doc := util.GoQueryDocWithCookies(http_url, cookies, timeout)
		reg := regexp.MustCompile("(\\s{2,})|(\\n)")
		doc.Find(".searchItem").Each(func(i int, s *goquery.Selection) {
			aTag := s.Find(".searchItemTitle").Find("a").First()
			href,_ := aTag.Attr("href")
			title := reg.ReplaceAllString(aTag.Text()," ")
			desc := reg.ReplaceAllString(s.Find(".searchCon").Text()," ")
			info := reg.ReplaceAllString(s.Find(".searchItemInfo").First().Text()," ")
			result := entity.ResultInfo{
				From:"http://blog.csdn.net", Title:title, Info:info,
				Href:href, Desc:strings.Replace(desc, "\"", "'", -1)}
			arr = append(arr, result)
		})

	}
	chan_result <- arr
}


