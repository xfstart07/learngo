// Author: xufei
// Date: 2020/10/23

package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MySqlHandler *gorm.DB

func InitMySql() {
	open, err := gorm.Open("mysql", "root:123456@tcp(local.bjywkd.test:3306)/learngo?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}

	open.LogMode(true)
	open.DB().SetMaxIdleConns(10)
	open.DB().SetMaxOpenConns(100)
	open.DB().SetConnMaxLifetime(7200)

	MySqlHandler = open
}
