package main

import (
	"goPerson/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


var router = gin.Default()


func main() {
	router.GET("/", index)
	router.GET("/person", handlers.GetPerson)
	router.GET("/person/:name", handlers.GetSinglePerson)
	router.POST("/person", handlers.PostPerson)
	router.PUT("/person/:name", handlers.UpdatePerson)
	router.DELETE("/person/:name", handlers.DeletePerson)
	router.GET("/founder", founder)
	log.Fatal(router.Run(":8080"))
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted,"You have reached the Deep State");
}

func founder(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Deep state is founded by Marcus")
}

