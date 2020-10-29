// Author: xufei
// Date: 2020/10/27

package main

import (
	"math/rand"

	db2 "learngo/pkg/gorm-practice/db"
	"learngo/pkg/gorm-practice/model"

	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
)

func main() {
	db2.InitMySql()

	db := db2.MySqlHandler
	db.AutoMigrate(model.UserInfo{})

	for i := 0; i < 2; i++ {
		send(db)
	}
}

func send(db *gorm.DB) {
	tx := db.Begin()
	for i := 0; i < 10000; i++ {
		age := rand.Intn(45) + 1
		tx.Save(&model.UserInfo{
			Name: faker.Name(),
			Age:  age,
		})
	}
	tx.Commit()
}
