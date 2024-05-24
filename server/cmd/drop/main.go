package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"my-app/db"
	"my-app/types"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.CreateDatabase(); err != nil {
		log.Fatal(err)
	}

	clearTable(db.DB, &types.User{})

	fmt.Println("Database cleared successfully.")
}

func clearTable(db *gorm.DB, model interface{}) {
	if err := db.Unscoped().Where("1 = 1").Delete(model).Error; err != nil {
		panic(err)
	}
}
