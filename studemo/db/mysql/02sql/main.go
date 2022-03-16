package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Connect("mysql", "dbapp:123456@tcp(127.0.0.1:3306)/datas?charset=utf8")
	if err != nil {
		fmt.Println("open db error:", err)
	}

	rows, err := db.Query("SELECT 'country', 'city', 11 as telcode FROM dual")
	if err != nil {
		fmt.Println("db query error:", err)
		return
	}
	// iterate over each row
	for rows.Next() {
		var country string
		// note that city can be NULL, so we use the NullString type
		var city sql.NullString
		var telcode int
		err = rows.Scan(&country, &city, &telcode)
		if err != nil {
			fmt.Println("db query error:", err)
			return
		}
		fmt.Println(country, city, telcode)

	}
	// check the error from rows
	// err = rows.Err()

	db.Close()
}
