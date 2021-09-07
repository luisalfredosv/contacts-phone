package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {

	gorm.Model
	
	Name string `json:"name" validate:"required,min=4,max=15"`
	Age uint `json:"age" validate:"required,min=10,max=120"`
	Phone string `json:"phone" gorm:"size:20"`
	Address string `json:"address"`
	Email string `json:"email"`
	Description string `json:"description" gorm:"type:TEXT"`

}