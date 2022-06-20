package main

import (
	"learngo/thirdparty/gorm-practice/db"
	"learngo/thirdparty/gorm-practice/model"
	"log"
)

func main() {
	db.InitSqlite()

	if err := db.SqliteHandler.AutoMigrate(&model.UserInfo{}, &model.Company{}); err != nil {
		panic(err)
	}

	//mock data
	c := &model.Company{
		Name: "weixin",
	}
	db.SqliteHandler.Create(c)
	user := model.UserInfo{
		Name:      "userxx",
		Age:       10,
		CompanyID: c.ID,
	}
	if err := db.SqliteHandler.Create(&user).Error; err != nil {
		panic(err)
	}

	var users []model.UserInfo
	if err := db.SqliteHandler.Preload("Company").Find(&users).Error; err != nil {
		panic(err)
	}

	log.Printf("users: %+v", users)

	db.CleanSqlite()
}
