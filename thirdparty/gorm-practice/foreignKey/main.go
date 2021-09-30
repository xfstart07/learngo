// Author: xufei
// Date: 2020/10/23

package main

import (
	"fmt"

	db2 "learngo/thirdparty/gorm-practice/db"
	"learngo/thirdparty/gorm-practice/model"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

func main() {
	db2.InitMySql()

	db := db2.MySqlHandler
	db.AutoMigrate(model.Class{}, model.Student{}, model.StuCard{})

	//send(db)
	//sendStuCard(db)

	// Note: Has Many 使用 Preload
	var class model.Class
	db.Model(model.Class{}).Preload("Students").Where("id = 5").Find(&class)
	fmt.Printf("%+v\n", class)

	// Note: Has one 和 Belongs To，使用 Preload
	var stu model.Student
	if err := db.Model(model.Student{}).Preload("Class").Preload("StuCard").First(&stu).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", stu)
}

func send(db *gorm.DB) {

	clss1 := model.Class{Name: "一班"}
	db.Create(&clss1)

	clss2 := model.Class{Name: "二班"}
	db.Create(&clss2)

	for i := 0; i < 10; i++ {
		stu := model.Student{
			Name:    fmt.Sprintf("学生%d", i),
			ClassId: clss1.ID,
		}
		db.Create(&stu)
	}

	for i := 0; i < 10; i++ {
		stu := model.Student{
			Name:    fmt.Sprintf("二学生%d", i),
			ClassId: clss2.ID,
		}
		db.Create(&stu)
	}
}

func sendStuCard(db *gorm.DB) {
	var stus []model.Student
	db.Model(model.Student{}).Find(&stus)

	for idx := range stus {
		stus[idx].StuCard.Number = xid.New().String()
		if err := db.Save(&stus[idx]).Error; err != nil {
			fmt.Println(err)
		}
	}
}
