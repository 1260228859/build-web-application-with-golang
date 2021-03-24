package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:'2017@ebola'@bj-cynosdbmysql-grp-io3qyr86.sql.tencentcdb.com:27816/yibaidb?charset=utf8mb4")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	fmt.Println(SqlDB)
	if err != nil {
		log.Fatal(err.Error())
	}
}
