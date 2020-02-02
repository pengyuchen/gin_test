package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works.")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	log.Println(firstName)
	log.Println(lastName)

	p := models.Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func AddMultiPersonApi(c *gin.Context) {
	var persons models.Persons
	// var persons []models.Person
	err := c.BindJSON(&persons)

	if err != nil {
		log.Fatalln(err)
	}
	for _, person := range persons.Persons {

		firstName := person.FirstName
		lastName := person.LastName

		p := models.Person{FirstName: firstName, LastName: lastName}

		ra, err := p.AddPerson()
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

func DelPersonsIdApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("id"))

	log.Println(id)

	p := models.Person{Id: id}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.Person{Id: id}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.Person{Id: id}

	err := c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	p := models.Person{Id: id}

	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func GetPersonsApi(c *gin.Context) {
	p := models.Person{}

	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
