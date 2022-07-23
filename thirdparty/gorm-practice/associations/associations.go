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

	if err := db.SqliteHandler.AutoMigrate(&model.UserInfo{}, &model.Language{}, &model.Order{}, &model.OrderItem{}, &model.Product{}); err != nil {
		panic(err)
	}

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

	//	PreloadAll
	preloadAll()
}

func preloadAll() {
	var user model.UserInfo
	if err := db.SqliteHandler.First(&user).Error; err != nil {
		log.Fatal(err)
	}

	pid := model.Product{Name: "衣服"}
	db.SqliteHandler.Create(&pid)

	order := model.Order{
		Price:  100,
		UserID: user.ID,
		OrderItem: model.OrderItem{
			ItemName:  "订单额外信息",
			ProductID: pid.ID,
		},
	}
	db.SqliteHandler.Create(&order)

	var users []model.UserInfo
	db.SqliteHandler.Preload("Orders.OrderItem.Product").Find(&users)
	//SELECT * FROM `product` WHERE `product`.`id` = 1
	//SELECT * FROM `orderItem` WHERE `orderItem`.`order_id` = 1
	//SELECT * FROM `order` WHERE `order`.`user_id` = 1
	//SELECT * FROM `user`
	log.Printf("%+v", users)
}
