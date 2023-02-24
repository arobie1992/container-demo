package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/greet", greetHandler)
	r.Run(":8086")
}

func greetHandler(c *gin.Context) {
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)

	var json HttpReq
	if err := c.ShouldBindJSON(&json); err != nil {
		logger.Printf("Error occurred while parsing json %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Printf("Received request to greet %s\n", json.Name)
	c.JSON(http.StatusOK, HttpResp{"Hi, " + json.Name + "!"})
}

type HttpReq struct {
	Name string `json:"name"`
}

type HttpResp struct {
	Message string `json:"message"`
}
