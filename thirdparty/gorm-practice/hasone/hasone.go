package main

import (
	"learngo/thirdparty/gorm-practice/db"
	"learngo/thirdparty/gorm-practice/model"
	"log"
)

func main() {
	db.InitSqlite()
	defer db.CleanSqlite()

	if err := db.SqliteHandler.AutoMigrate(&model.UserInfo{}, &model.CreditCard{}); err != nil {
		panic(err)
	}

	user := model.UserInfo{
		Name:      "userxx",
		Age:       10,
		CompanyID: 1,
	}
	if err := db.SqliteHandler.Create(&user).Error; err != nil {
		panic(err)
	}

	cred := model.CreditCard{
		Number: "123123",
		UserID: user.ID,
	}
	db.SqliteHandler.Create(&cred)

	var u model.UserInfo
	db.SqliteHandler.Preload("CreditCard").First(&u)
	// SELECT * FROM `creditcard` WHERE `creditcard`.`user_id` = 1
	// SELECT * FROM `user` ORDER BY `user`.`id` LIMIT 1
	log.Printf("user %+v\n", u)

}
