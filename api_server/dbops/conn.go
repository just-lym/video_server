package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123!@#@tcp(localhost:3306)/video_server?charset=utf-8")
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(3)
	dbConn.SetMaxIdleConns(int(time.Second * 10))
	if err != nil {
		panic(err.Error())
	}
}
