package db

import (
	"log"
	"rest-api-tokobuku-echo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	conf := config.GetConfig()

	connection := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	database, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}
	db = database

}

func CreateCon() *gorm.DB {
	return db
}
