package config

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var db_detail = "root:@tcp(127.0.0.1:3306)/pos_bangkit?parseTime=true"

	db, err := sql.Open("mysql", db_detail)

	if err != nil {
		fmt.Println("Database tidak konek")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Database tidak konek")
		panic(err)
	}

	fmt.Println("Database Konek")
	DB = db
}