// Author: xufei
// Date: 2020/7/20

package main

import (
	"log"

	"learngo/thirdparty/gorm-practice/model"

	"github.com/bxcodec/faker/v3"
)

func main() {
	db := model.GetDB()

	for i := 0; i < 100000; i++ {
		var fake model.User
		err := faker.FakeData(&fake)
		if err != nil {
			log.Fatal(err)
		}

		user := model.User{
			UUID:      fake.UUID,
			Name:      fake.Name,
			Email:     fake.Email,
			Cellphone: fake.Cellphone,
			Gender:    fake.Gender,
			Password:  fake.Password,
			Address:   fake.Address,
			Avatar:    fake.Avatar,
		}

		if err := db.Save(&user).Error; err != nil {
			log.Fatal(err)
		}
	}
}
