package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/go-redis/redis"

	logger "github.com/sirupsen/logrus"
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
	raw, _ := json.Marshal(res)
	fmt.Fprintf(w, string(raw))
}

func httpResponseData(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "application/json;charset:utf-8;")
	fmt.Fprintf(w, data)
}

//接口请求
type RequestInfo struct {
	IName       string                 `json:"Iname"`       // 接口名称
	QueryParams map[string]interface{} `json:"QueryParams"` // 接口查询类型 NameQuery Query
}

//计算请求信息生成唯一性参数
func getRequestId(reqparams map[string]interface{}) string {
	b, err := json.Marshal(reqparams)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return redis.Nil.Error()
	}

	h := md5.New()
	h.Write([]byte(b))
	return hex.EncodeToString(h.Sum(nil))
}

func rows2map(rows2 *sqlx.Rows) map[int]map[string]string {
	//返回所有列
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k := range vals {
		scans[k] = &vals[k]
	}

	//将所有结果封装到一个map中返回key为index的值
	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			// fmt.Printf(string(v))
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	// fmt.Println(result1)
	// for k, v := range result1 {
	// 	fmt.Println("第", k, "行", "===>", v)
	// }
	return result
}

func db_namequery(select_sql string, params interface{}) (map[int]map[string]string, error) {
	rows2, errs := db.NamedQuery(select_sql, params)
	if errs != nil {
		fmt.Println("open db error:", errs)
		return nil, errs
	}
	defer rows2.Close()
	return rows2map(rows2), nil
}

func db_query(select_sql string) (map[int]map[string]string, error) {
	rows3, errs := db.Queryx(select_sql)
	if errs != nil {
		fmt.Println("open db error:", errs)
		return nil, errs
	}
	defer rows3.Close()
	return rows2map(rows3), nil
}

type IdataConf struct {
	Name        string `db:"it_name"`      // 接口名称
	QueryType   string `db:"query_type"`   // 接口查询类型 NameQuery Query
	RowType     int    `db:"row_type"`     // 返回数据是条数
	CacheTime   int    `db:"cache_time"`   //缓存时间
	QueryString string `db:"query_string"` //查询字符串
}

func init_db() {
	var err error
	db, err = sqlx.Connect("mysql", "dbapp:123456@tcp(127.0.0.1:3306)/datas?charset=utf8")
	if err != nil {
		logger.Error("lopen db error:", err)
		return
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(40)
}

func load_querylist() {
	select_sql := "select it_name ,query_type ,row_type ,cache_time ,query_string  from app_table_httpquery where f_status=1"
	rows, err := db.Queryx(select_sql)
	if err != nil {
		logger.Error("load queryList error:", err)
		return
	}

	result := make(map[string]IdataConf)
	for rows.Next() {
		idata := IdataConf{}
		err := rows.StructScan(&idata)
		if err != nil {
			logger.Fatalln(err)
		}
		key := idata.Name
		result[key] = idata

	}
	query_list = result
}

func init_redis() {
	rediscli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		PoolSize: 100,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := rediscli.Ping().Result()
	logger.Info("init redis :", pong, err)

}

func setRedis(redisKey string, data interface{}, tt time.Duration) {

	tmp, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	status := rediscli.Set(redisKey, string(tmp), tt)
	logger.Debug("set redis > key:", redisKey, "expire_time:", tt, "status:", status.Val())
}

///////////////////////////////////////////////////////////////////////
// http响应部分
func http_query(w http.ResponseWriter, r *http.Request) {
	var req RequestInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := InstanceData(http.StatusBadRequest, err.Error(), nil)
		httpResponseCode(w, res)
	} else {
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
		res, _ := rediscli.Get(redisKey).Result()
		if len(res) > 0 {
			httpResponseData(w, res)
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
		}

		if it_item.RowType > 1 {
			res := InstanceData(http.StatusOK, "success", data)
			setRedis(redisKey, res, expire_time)
			httpResponseCode(w, res)
			return
		} else {
			if tt, ok := data[0]; ok {
				res := InstanceData(http.StatusOK, "success", tt)
				setRedis(redisKey, res, expire_time)
				httpResponseCode(w, res)
				return

			} else {
				res := InstanceData(http.StatusOK, "success", nil)
				setRedis(redisKey, res, expire_time)
				httpResponseCode(w, res)
				return
			}
		}
	}
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
var query_list map[string]IdataConf
var rediscli *redis.Client

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
	init_db()
	load_querylist()
	init_redis()
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
