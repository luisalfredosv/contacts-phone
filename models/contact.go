package models

import "github.com/jinzhu/gorm"

type Contact struct {

	gorm.Model
	
	Name string `json:"name"`
	Age uint `json:"age"`
	phone string `json:"phone" gorm:"size:20"`
	address string `json:"address"`
	email string `json:"email"`
	description string `json:"description" gorm:"type:TEXT"`

}