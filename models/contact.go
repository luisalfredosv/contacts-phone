package models

import "github.com/jinzhu/gorm"

type Contact struct {

	gorm.Model
	
	Name string `json:"name"`
	Age uint `json:"age"`
	Phone string `json:"phone" gorm:"size:20"`
	Address string `json:"address"`
	Email string `json:"email"`
	Description string `json:"description" gorm:"type:TEXT"`

}