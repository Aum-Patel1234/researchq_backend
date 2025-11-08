package main

import (
	"github.com/Aum-Patel1234/researchq_backend/initializers"
	"github.com/Aum-Patel1234/researchq_backend/models"
)

func main() {
	initializers.LoadEnv()
	db := initializers.ConnectToDB()

	db.AutoMigrate(&models.User{})
}
