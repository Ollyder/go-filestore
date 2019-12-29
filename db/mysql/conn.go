package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/filestore?charset=utf8")
	db.SetMaxOpenConns(100)
	err := db.Ping()
	if err != nil {
		fmt.Printf("Connect to db err : %s \n", err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
