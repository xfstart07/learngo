package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqliteHandler *gorm.DB

func InitSqlite() {
	dbpath := "./gormsqlite.db?_timeout=5000"
	open, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		panic(err)
	}

	SqliteHandler = open.Debug()
}

func CleanSqlite() {
	os.Remove("./gormsqlite.db")
}
