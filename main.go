package main

import (
	"net/http"

	"./routers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "gopkg.in/gin-gonic/gin.v1"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})
	routers.InitRouter()

	router.Run(":8000")
}
