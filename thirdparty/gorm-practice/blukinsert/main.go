// Author: xufei
// Date: 2020/10/27

package main

import (
	"fmt"
	"math/rand"
	"time"

	db2 "learngo/thirdparty/gorm-practice/db"
	"learngo/thirdparty/gorm-practice/model"

	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
)

func main() {
	db2.InitMySql()

	db := db2.MySqlHandler
	db.AutoMigrate(model.UserInfo{})

	t := time.Now()
	send(db)
	fmt.Println(time.Since(t))
}

func send(db *gorm.DB) {
	var list []interface{}
	for i := 0; i < 300000; i++ {
		age := rand.Intn(45) + 1
		info := model.UserInfo{
			Name: faker.Name(),
			Age:  age,
		}

		list = append(list, info)
	}

	err := gormbulk.BulkInsert(db, list, 3000)
	if err != nil {
		panic(err)
	}
}
