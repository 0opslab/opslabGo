package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
	rows2.Close()
	return result
}
func main() {
	var err error
	var db *sqlx.DB
	db, err = sqlx.Connect("mysql", "dbapp:123456@tcp(127.0.0.1:3306)/datas?charset=utf8")
	if err != nil {
		fmt.Println("open db error:", err)
	}

	rows2, errs := db.NamedQuery(`SELECT 'country', 'city', 11 as telcode FROM dual where 1=:fn`,
		map[string]interface{}{"fn": 1, "fn2": 2})
	if errs != nil {
		fmt.Println("open db error:", errs)
		return
	}
	result := rows2map(rows2)

	fmt.Println(result)
	for k, v := range result {
		fmt.Println("第", k, "行", "===>", v)
	}

	rows3, errs := db.Queryx(`select * from t_stock_info limit 2`)
	if errs != nil {
		fmt.Println("open db error:", errs)
		return
	}
	result1 := rows2map(rows3)

	fmt.Println(result1)
	for k, v := range result1 {
		fmt.Println("第", k, "行", "===>", v)
	}

	db.Close()
}
