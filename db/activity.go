package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Activity 文章列表
type Activity struct {
	gorm.Model
	Title   string    `gorm:"type:varchar(255)" json:"title" binding:"required"`
	Cover   string    `gorm:"type:varchar" json:"cover" binding:"required"`
	Content string    `gorm:"type:text" json:"content" binding:"required"`
	StartAt time.Time `gorm:"type:timestamp" json:"start_at" bind:"required"`
	EndAt   time.Time `gorm:"type:timestamp" json:"end_at" bind:"required"`
	UserID  uint
}
