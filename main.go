// package main

// import (
// 	db "./database"
// 	"./routers"
// )

// func main() {
// 	defer db.SqlDB.Close()
// 	r := routers.InitRouter()
// 	r.Run(":8080")
// }

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "gopkg.in/gin-gonic/gin.v1"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.Run(":8000")
}
