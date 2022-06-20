// Author: xufei
// Date: 2020/10/23

package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlHandler *gorm.DB

func InitMySql() {
	dns := "root:123456@tcp(local.bjywkd.test:3306)/learngo?charset=utf8mb4&loc=Local&parseTime=true"
	open, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	MySqlHandler = open
}
