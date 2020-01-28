// package database

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var SqlDB *sql.DB

// func init() {

// 	db, err := SqlDB.Open("mysql", "jr:12345678@tcp(127.0.0.1:3306)/test?parseTime=true")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer db.Close()

// 	db.SetMaxIdleConns(20)
// 	db.SetMaxOpenConns(20)

// 	if err := db.Ping(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql",
		"jr:12345678@tcp(127.0.0.1:3306)/gin_exer?parseTime=true&loc=Local&charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
