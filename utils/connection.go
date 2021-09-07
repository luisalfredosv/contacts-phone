package utils

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "userDb:pwdDb@/dbName?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	return db
}