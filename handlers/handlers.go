package handlers

import (
	"goPerson/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var persons[]models.Person = []models.Person{{Name: "Marcus", 
											  Age: 32,
											  City: "Milan",
											  Country: "Italy" }}
// GetPerson returns all person ...
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
// UpdatePerson updates person ...
func UpdatePerson(c *gin.Context) {
	var name = c.Param("name")
	for i, p := range persons {
		if p.Name == name {
			var updatePerson models.Person
			if err := c.ShouldBindJSON(&updatePerson); err != nil {
				c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
				return
			}
			persons[i] = updatePerson
			c.JSON(http.StatusAccepted, persons)
			return
		}
	}
	c.JSON(http.StatusNotFound, "No person with name " + name)
}

// DeletePerson deletes a person ...
func DeletePerson(c *gin.Context) {
	var name = c.Param("name")
	for i, p := range persons {
		if p.Name == name {
			persons = append(persons[:i],persons[i+1:] ...)
			c.JSON(http.StatusAccepted, persons)
			return
		}
	}
	c.JSON(http.StatusNotFound, "No person with name " + name)
}
// GetSinglePerson returns a person ...
func GetSinglePerson(c *gin.Context) {
	var name = c.Param("name")
	for _, p := range persons {
		if p.Name == name {
			c.JSON(http.StatusAccepted, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, "No person with name " + name)
}