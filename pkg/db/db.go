package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func MySQLConnect() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	DB, err = gorm.Open("mysql", Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("DB Connection Error: ")
	} else {
		log.Println("DB Connection Success")
	}

	DBMigrate()

}

func DBMigrate() {}
