package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(weixinote.dev:3306)/learngo?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(true)

	db.AutoMigrate(&User{})

	return db
}
