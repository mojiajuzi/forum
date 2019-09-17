package model

import (
	"fmt"

	"github.com/mojiajuzi/forum/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	name := config.Config("DB_USER", "forum")
	pass := config.Config("DB_PASSWORD", "")
	host := config.Config("DB_HOST", "localhost")
	port := config.Config("DB_PORNT", "3306")
	dnName := config.Config("DB_NAME", "forum")
	char := config.Config("DB_CHARSET", "forum")
	parsetTime := config.Config("DB_PARSET_TIME", "forum")
	loc := config.Config("DB_LOC", "forum")
	connect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", name, pass, host, port, dnName, char, parsetTime, loc)
	dbConnect, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("数据连接异常")
	}
	db = dbConnect
}

//Db 获取数据库连接
func Db() *gorm.DB {
	return db
}

//Migrate 数据迁移
func Migrate() {
	fmt.Println(db)
	userMigrate()
	activityMigrate()
}

//用户表
func userMigrate() {
	db.CreateTable(&User{})
}

//活动表
func activityMigrate() {
	db.CreateTable(&Activity{})
}
