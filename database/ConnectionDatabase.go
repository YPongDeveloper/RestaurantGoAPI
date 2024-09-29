// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB, err = gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

func Connect() (*gorm.DB, error) {
	dsn := "user:password@tcp(localhost:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Database connected successfully.")
	DB = db
	return db, nil
}
