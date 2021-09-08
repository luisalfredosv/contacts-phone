package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {

	gorm.Model
	
	Name string `json:"name" validate:"required"`
	Age uint `json:"age" validate:"required,gte=0,lte=130"`
	Phone string `json:"phone" gorm:"size:20" validate:"required"`
	Address string `json:"address" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Description string `json:"description" gorm:"type:TEXT"`

}