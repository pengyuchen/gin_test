package routers

import (
	"../apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", apis.IndexApi)
	r.POST("/person", apis.AddPersonApi)
	r.DELETE("/person/:id", apis.DelPersonApi)
	r.PUT("/person/:id", apis.ModPersonApi)
	r.GET("/person/:id", apis.GetPersonApi)
	r.GET("/persons", apis.GetPersonsApi)

	r.POST("/persons/DeleteAllById", apis.DelPersonsIdApi)
	r.POST("/persons/AddMultiPersons", apis.AddMultiPersonApi)

	return r
}
