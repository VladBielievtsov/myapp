package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func CreateDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}
	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		dbport     = "5432"
		uri        = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbhost, dbuser, dbpassword, dbname, dbport)
	)
	DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
	}
	log.Println("🚀 Connected Successfully to the Database")
	return nil
}
