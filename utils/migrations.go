package utils

import (
	"fmt"
	"github.com/luisalfredosv/golang-gorilla/models"
)

func MigrateDB(){
	db := GetConnection()
	defer db.Close()
	
	fmt.Println("Migrating models...")

	db.AutoMigrate(&models.Contact{})
}