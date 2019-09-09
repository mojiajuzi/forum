package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Activity 文章列表
type Activity struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255)" json:"title" validate:"required"`
	Cover       string    `gorm:"type:varchar(200)" json:"cover" validate:"required"`
	Content     string    `gorm:"type:text" json:"content" validate:"required"`
	StartAt     time.Time `json:"start_at" validate:"required"`
	EndAt       time.Time `json:"end_at" validate:"required"`
	WillTotal   uint      `gorm:"type:int" json:"will_total" validate:"will_total"`
	ActualTotal uint      `gorm:"type:int" json:"actual_total"`
	UserID      uint
}
