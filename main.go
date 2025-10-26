package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	r := gin.Default()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.GET("/", func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Welcome to home..")
		c.String(http.StatusOK, "Welcome to Home..")
	})

	r.Run(os.Getenv("PORT"))
}
