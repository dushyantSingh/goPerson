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
	router.DELETE("/person/:name", founder)
	router.GET("/person", handlers.GetPerson)
	router.POST("/person", handlers.PostPerson)
	router.GET("/founder", founder)
	log.Fatal(router.Run(":8080"))
}

func index(c *gin.Context) {
	c.JSON(http.StatusAccepted,"You have reached the Deep State");
}

func founder(c *gin.Context) {
	c.JSON(http.StatusAccepted, "Deep state is founded by Marcus")
}

