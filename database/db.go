package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "sigolang-sv/configs"
)

var (
	DB *gorm.DB
	err error
)

func MySQLConnect() {
	dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("DB Connection Error: ")
	} else {
		log.Println("DB Connection Success")
	}

	DBMigrate()

}