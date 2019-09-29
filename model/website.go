package model

import "github.com/jinzhu/gorm"

//Website 网站类型
type Website struct {
	gorm.Model
	Logo    string `gorm:"not null;unique"`
	URL     string `gorm:"not null;unique"`
	Score   uint
	Against uint
	Agree   uint
	Status  uint
	UserID  uint
}
