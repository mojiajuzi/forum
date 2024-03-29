package model

import (
	"github.com/jinzhu/gorm"
)

//User 用户表
type User struct {
	gorm.Model
	Status   int8       `gorm:"default:1;type:int" json:"status"`
	Avatar   string     `gorm:"type:varchar(200)" json:"avatar"`
	Email    string     `gorm:"type:varchar(100);unique_index" validate:"required,email" json:"email"`
	Name     string     `gorm:"type:varchar(50)" validate:"required,min=6,max=20" json:"name"`
	Password string     `gorm:"type:varchar(255)" validate:"required,min=6,max=20" json:"password"`
	Activity []Activity `json:"activity"`
}
