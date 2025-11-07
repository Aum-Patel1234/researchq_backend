package main

import (
	"log"
	"os"
	"strings"

	"github.com/Aum-Patel1234/researchq_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load .env — continuing with environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	// fmt.Println("Starting server on", port)

	r := routes.SetUpRoutes()

	originsEnv := os.Getenv("ALLOWED_ORIGINS")
	var allowedOrigins []string
	for _, origin := range strings.Split(originsEnv, ",") {
		allowedOrigins = append(allowedOrigins, origin)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
	}))
	// TODO: in future configure this if needed
	// config := cors.Config{
	// 	AllowOrigins: []string{"*"}, // or specify e.g. []string{"http://localhost:3000"}
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }
	// r.Use(cors.New(config))

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server : %v", err)
	}
}
