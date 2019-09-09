package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mojiajuzi/forum/service"
)

var db *gorm.DB

func init() {
	fmt.Println("数据库初始化")
	name := service.Config("DB_USER", "forum")
	pass := service.Config("DB_PASSWORD", "")
	host := service.Config("DB_HOST", "localhost")
	port := service.Config("DB_PORNT", "3306")
	dnName := service.Config("DB_NAME", "forum")
	char := service.Config("DB_CHARSET", "forum")
	parsetTime := service.Config("DB_PARSET_TIME", "forum")
	loc := service.Config("DB_LOC", "forum")
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
