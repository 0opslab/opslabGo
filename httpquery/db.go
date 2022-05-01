package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	logger "github.com/sirupsen/logrus"
)

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

func init_db(mysql_info string) {
	var err error
	db, err = sqlx.Connect("mysql", mysql_info)
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
