package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息
func InstanceData(code int, message string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  message,
		Data: data,
	}
}

//统一响应
func httpResponseCode(w http.ResponseWriter, res Response) {
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	fmt.Fprint(w, toJosn(res))
}

func httpResponseJson(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	fmt.Fprint(w, data)
}

//接口请求
type RequestInfo struct {
	IName       string                 `json:"Iname"`       // 接口名称
	QueryParams map[string]interface{} `json:"QueryParams"` // 接口查询类型 NameQuery Query
}

//计算请求信息生成唯一性参数
func getRequestId(reqparams map[string]interface{}) string {
	request := toJosn(reqparams)
	h := md5.New()
	h.Write([]byte(request))
	return hex.EncodeToString(h.Sum(nil))
}

var query_list map[string]IdataConf
