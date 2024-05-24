package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my-app/types"
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

	var count int64
	DB.Model(&types.User{}).Count(&count)

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		id := uuid.New()
		defaultUser := types.User{
			ID:       &id,
			Username: "admin",
			Password: string(hashedPassword),
		}
		result := DB.FirstOrCreate(&types.User{}, defaultUser)
		if result.Error != nil {
			panic(result.Error)
		}
		fmt.Println("Default user created successfully!")
		fmt.Println("====================")
		fmt.Println("Username: admin")
		fmt.Println("Password: password")
		fmt.Println("====================")
	}

	return nil
}

func Migrate() {
	err := DB.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}
