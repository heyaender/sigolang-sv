package database

import (
	"fmt"
	"log"
	"sigolang-sv/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func MySQLConnect() {
	dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", configs.DB_USER, configs.DB_PASSWORD, configs.DB_NAME)
	// dsn := ("andri:j26dlZaoZ8wQpxT8@tcp(127.0.1:3306)/andri?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Println("DB Connection Error: ")
	} else {
		log.Println("DB Connection Success")
	}

	DBMigrate()

}

