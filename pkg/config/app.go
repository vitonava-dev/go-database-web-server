package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	fmt.Println("Connecting to database")
	dsn := "root:navabi05@tcp(127.0.0.1:3306)/book_store?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db = d
	fmt.Println("Database connected")
}

func GetDB() *gorm.DB {
	return db
}
