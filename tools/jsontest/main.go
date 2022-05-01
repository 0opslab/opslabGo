package main

import (
	"encoding/json"
	"fmt"
)

func parseJson(b []byte, data interface{}) error {
	err := json.Unmarshal(b, &data)
	return err
}

func parseJsonString(b string, data interface{}) error {
	err := json.Unmarshal([]byte(b), &data)
	return err
}

func toJosn(data interface{}) string {
	tt, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	return string(tt)
}

//接口请求
type RequestInfo struct {
	IName       string                 `json:"Iname"`       // 接口名称
	QueryParams map[string]interface{} `json:"QueryParams"` // 接口查询类型 NameQuery Query
}

func main() {

	jsonss := "{\"Iname\": \"test_1\",\"QueryParams\": {\"fn\":1,\"fn2\":2}}"
	var req RequestInfo

	// err := parseJson([]byte(jsonss), &req)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := parseJsonString(jsonss, &req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.IName)
	fmt.Println(req.QueryParams)
	fmt.Println(string(toJosn(req)))
}
