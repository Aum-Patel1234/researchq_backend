package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/Aum-Patel1234/researchq_backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("DSN not found in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get generic database object: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("❌ Database not reachable: %v", err)
	}

	if utils.IsDev() {
		fmt.Println("✅ Connected to PostgreSQL successfully!")
	}
	return db
}
