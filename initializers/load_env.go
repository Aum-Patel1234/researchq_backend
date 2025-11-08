package initializers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load .env — continuing with environment variables")
	}

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
}
