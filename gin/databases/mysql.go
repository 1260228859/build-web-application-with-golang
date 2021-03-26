package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var sqlDB *sql.DB

func init() {
	var err error
	sqlDB, err = sql.Open("mysql", "root:'2017@ebola'@bj-cynosdbmysql-grp-io3qyr86.sql.tencentcdb.com:27816/ivms-dev1?charset=utf8mb4")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
