package main

import (
	"fmt"

	"learngo/pkg/gorm-practice/model"
)

func main() {
	db := model.GetDB()
	fmt.Println(db)

	user := model.User{}
	if err := db.Where(model.User{Cellphone: "12345678"}).Assign(model.User{Name: "xufei"}).FirstOrCreate(&user).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
