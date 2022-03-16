package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jmoiron/sqlx"
)

// option mysql
// 多看文档 https://pkg.go.dev/github.com/jmoiron/sqlx
var Db *sql.DB

type Message struct {
	Message string `db:"message"`
}

func init_db() {
	var err error
	Db, err = sql.Open("mysql", "dbapp:123456@tcp(127.0.0.1:3306)/datas?charset=utf8")
	if err != nil {

		fmt.Println("conn failed!")
		log.Fatal(err.Error())
	}
	//defer Db.Close()

	fmt.Println(&Db)

	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)

	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("db inti")
}
func main() {
	init_db()

	rows, err1 := Db.Query("select 1 as message,'test' from dual")
	if err1 != nil {
		fmt.Println("select error, ", err1)
		return
	}

	for rows.Next() {
		var id int
		var post string
		err := rows.Scan(&id, &post)
		if err != nil {
			fmt.Println("read-error, ", err)
			return
		}

		fmt.Println(id)
		fmt.Println(post)

	}
	rows.Close()

}
