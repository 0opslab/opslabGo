package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"time"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	logger "github.com/sirupsen/logrus"
)

// http响应部分
func http_query(w http.ResponseWriter, r *http.Request) {
	var req RequestInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := InstanceData(http.StatusBadRequest, err.Error(), nil)
		httpResponseCode(w, res)
		return
	}
	requestId := getRequestId(req.QueryParams)
	logger.Info("request:", req.IName, "[params]:", req.QueryParams, "requestId:", requestId)
	it_item, ok := query_list[req.IName]
	if !ok {
		res := InstanceData(http.StatusNotFound, "无该接口", nil)
		httpResponseCode(w, res)
		return
	}
	expire_time := time.Duration(it_item.CacheTime) * time.Second
	redisKey := req.IName + ":" + requestId
	res := getString(redisKey)
	if len(res) > 0 {
		httpResponseJson(w, res)
		return
	}

	var data map[int]map[string]string
	var sqlerr error
	if it_item.QueryType == "query" {
		data, sqlerr = db_query(it_item.QueryString)
	} else {
		data, sqlerr = db_namequery(it_item.QueryString, req.QueryParams)
	}
	if sqlerr != nil {
		res := InstanceData(http.StatusBadGateway, "Exception", nil)
		httpResponseCode(w, res)
		return
	}
	var ts interface{}
	if it_item.RowType > 1 {
		ts = data
	} else {
		if tt, ok := data[0]; ok {
			ts = tt
		} else {
			ts = nil
		}
	}
	result := InstanceData(http.StatusOK, "success", ts)
	setRedis(redisKey, result, expire_time)
	httpResponseCode(w, result)
}

//重新加载接口配置数据信息
func http_load_querylist(w http.ResponseWriter, r *http.Request) {
	var req RequestInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := InstanceData(http.StatusBadRequest, err.Error(), nil)
		httpResponseCode(w, res)
	} else {
		logger.Debug("request:", req.IName, "[params]:", req.QueryParams)
		load_querylist()
		res := InstanceData(http.StatusOK, "success", unsafe.Sizeof(query_list))
		httpResponseCode(w, res)
	}
}

//全局变量
var db *sqlx.DB

//程序入口及启动
func main() {

	flag.Parse()

	//lf, err := os.OpenFile(logPath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0660)
	//if err != nil {
	//	logger.Fatalf("Failed to open log file: %v", err)
	//}
	//defer lf.Close()
	//
	//defer logger.Init("LoggerExample", *verbose, true, lf).Close()
	logger.SetLevel(logger.InfoLevel)

	//设置访问的路由
	init_db("dbapp:123456@tcp(127.0.0.1:3306)/datas?charset=utf8")
	init_redis("127.0.0.1:6379")

	load_querylist()
	logger.Info("iload:", query_list)
	http.HandleFunc("/reload", http_load_querylist)
	http.HandleFunc("/data", http_query)

	//设置监听的端口
	errs := http.ListenAndServe(":10000", nil)
	if errs != nil {
		logger.Info("start error :", errs)
		return
	}
	logger.Info("service is started ... ")

}
