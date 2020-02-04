package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Gin_Mysql_api/models"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works.")
}

func AddUserApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	log.Println(firstName)
	log.Println(lastName)

	p := models.User{FirstName: firstName, LastName: lastName}

	ra, err := p.AddUser()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func AddUsersApi(c *gin.Context) {
	var users models.Users
	// var persons []models.Person
	err := c.BindJSON(&users)

	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users.Users {

		firstName := user.FirstName
		lastName := user.LastName

		p := models.User{FirstName: firstName, LastName: lastName}

		ra, err := p.AddUser()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("insert successful %d", ra)

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})

		// fmt.Printf("%+v\n", person.FirstName)   // show on terminal
		// c.JSON(http.StatusOK, person.FirstName) //show on postman
	}
}

func DelUserIdsApi(c *gin.Context) {
	var users models.Users
	// var persons []models.Person
	err := c.BindJSON(&users)

	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users.Users {

		id := user.Id

		log.Println(id)

		p := models.User{Id: id}

		ra, err := p.DelUser()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}
}

func DelUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}

	ra, err := p.DelUser()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func ModUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}

	err := c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := p.ModUser()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	p := models.User{Id: id}

	user, err := p.GetUser()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsersApi(c *gin.Context) {
	p := models.User{}

	users, err := p.GetUsers()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
