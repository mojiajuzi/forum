package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User 用户表
type User struct {
	gorm.Model
	Status   int8   `gorm:"default:1,type:tinyint(4)",json:"status"`
	Avatar   string `gorm:"",json:"avatar"`
	Email    string `gorm:"type:varchar(100);unique_index",validate:"required,email", json:"email"`
	Name     string `gorm:"type:varchar(50)",validate:"required,min=6,max=20", json:"name"`
	Password string `gorm:"type:varchar(255)",validate:"required,min=6,max=20"`
	Articles []Article
}

//Article 文章列表
type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255)",json:"title",validate:"required"`
	Cover   string `gorm:"type:varchar", json:"cover", validate:"required"`
	Content string `gorm:"type:text", json:"content", validate:"required"`
	Tag     []*Tag `gorm:"many2many:article_tags"`
	UserID  uint
}

//Tag 标签列表
type Tag struct {
	gorm.Model
	Articles []*Article `gorm:"many2many:article_tags"`
	Name     string     `gorm:"type:varchat(50)", json:"tag", validate:"required"`
}

//Comment 评论
type Comment struct {
	gorm.Model
	UserID    uint
	ArticleID uint
	Content   string `gorm:"type:text",json:"content",validate:"required"`
}
