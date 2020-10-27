// Author: xufei
// Date: 2020/10/26

package main

import (
	"fmt"

	db2 "learngo/pkg/gormexp/db"

	"github.com/jinzhu/gorm"
)

type Dog struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Toys []Toy  `gorm:"polymorphic:Owner"` // polymorphic 多态，Has many
}

type Toy struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	OwnerId   uint   `gorm:"not null"`
	OwnerType string `gorm:"not null"`
}

func main() {
	db2.InitMySql()

	db := db2.MySqlHandler
	db.AutoMigrate(Dog{}, Toy{})

	//send(db)

	var d Dog
	if err := db.Model(Dog{}).Preload("Toys").First(&d).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", d)

	var dk Dog
	db.Model(Dog{}).First(&dk)

	// 检验添加 多态对象
	// 更新Dog，并创建一个操作对象 Toys
	dk.Name = "dk"
	dk.Toys = append(dk.Toys, Toy{Name: "toy4"})
	if err := db.Save(&dk).Error; err != nil {
		fmt.Println(err)
	}
}

func send(db *gorm.DB) {
	db.Create(&Dog{Name: "tony", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
}
