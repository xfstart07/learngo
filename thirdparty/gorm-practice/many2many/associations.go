package main

import (
	"learngo/thirdparty/gorm-practice/db"
	"learngo/thirdparty/gorm-practice/model"
	"log"
)

//关联关系的 自动创建、更新
//https://gorm.io/zh_CN/docs/associations.html
func main() {
	db.InitSqlite()
	defer db.CleanSqlite()

	if err := db.SqliteHandler.AutoMigrate(&model.UserInfo{}, &model.Language{}); err != nil {
		panic(err)
	}
	//CREATE TABLE `user` (`id` integer,`name` text NOT NULL DEFAULT "",`age` integer NOT NULL DEFAULT 0,`company_id` integer NOT NULL,PRIMARY KEY (`id`))
	//CREATE TABLE `languages` (`id` integer,`name` text,PRIMARY KEY (`id`))
	//CREATE TABLE `user_languages` (`user_info_id` integer,`language_id` integer,PRIMARY KEY (`user_info_id`,`language_id`))

	user := model.UserInfo{
		Name: "jinzhu",
		Languages: []model.Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	if err := db.SqliteHandler.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

}
