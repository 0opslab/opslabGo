package entity

import (
	"strings"
	"regexp"
)



/**
 * 根据关键字在标题中搜索匹配关键字，然后计算连接的可用度
 */
func SetOrder(result *ResultInfo, key string) {
	/**
	 * 默认搜索权重
	 */
	FORM_ORDER := map[string]int{
		"https://www.baidu.com":7,
		"https://toutiao.io":8,
		"https://blog.csdn.net":6,
		"http://blog.csdn.net":7,
	}

	order := result.Order

	//根据搜索来源设置基础分值
	if v, ok := FORM_ORDER[result.From]; ok {
		order = v
	}
	//标题中包括关键字加5分
	if strings.Contains(result.Title, key) {
		order += 5
	}

	if strings.Contains(result.Desc, key) {
		order += 5
	}

	//描述中有关键字加5分
	result.Order = order

	//将数据字段进行过滤
	reg := regexp.MustCompile("(\\s{2,})|(\\\\)|(\")")
	result.Title = strings.Replace(result.Title,"\"","",-1)
	result.Desc = reg.ReplaceAllString(strings.Replace(result.Desc,"\"","",-1),"")
	result.Info = strings.Replace(result.Info,"\"","",-1)
}
