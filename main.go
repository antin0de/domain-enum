package main

import (
	"log"
	"os"

	"antin0.de/domain-enum/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file, ensure you have all environment variables set")
	}

	mysqlDsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	models.Migrate(db)
}
