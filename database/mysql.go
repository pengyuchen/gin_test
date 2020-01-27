package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err := sql.Open("mysql", "user:password@/dbname")
	// SqlDB, err = sql.Open("mysql",
	// 	"mysqlcli:12345678@tcp(10.68.7.24:3307)/gin_exer?parseTime=true&loc=Local&charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
