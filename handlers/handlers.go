package handlers

import (
	"goPerson/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var persons[]models.Person = []models.Person{{Name: "Marcus", Age: 32, City: "Milan", Country: "Italy"}}
// GetPerson returns person ...
func GetPerson(c *gin.Context) {
	c.JSON(http.StatusAccepted, persons)
}
// PostPerson saves person ...
func PostPerson(c *gin.Context) {
	var p models.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	if p.Age == 0 || p.Name == "" || p.City == "" {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	persons = append(persons, p)
	c.JSON(http.StatusAccepted, persons)
}