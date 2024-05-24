package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"my-app/db"
	"my-app/types"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.CreateDatabase(); err != nil {
		log.Fatal(err)
	}

	db.Migrate()

	tx := db.DB.Begin()
	seedUsers(10000, tx)
	tx.Commit()
}

func seedUsers(quantity int, db *gorm.DB) {
	start := time.Now()
	for i := 0; i < quantity; i++ {

		id := uuid.New()
		user := types.User{
			ID:       &id,
			Username: id.String(),
			Password: "password",
		}

		if err := db.Create(&user).Error; err != nil {
			log.Printf("Could not create users %d: %v", i+1, err)
			continue
		}
	}
	duration := time.Since(start)
	msg := fmt.Sprintf("Users have been created in %s", duration)
	fmt.Println(msg)
}
