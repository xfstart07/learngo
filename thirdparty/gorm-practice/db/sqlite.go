package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqliteHandler *gorm.DB

func InitSqlite() {
	dbpath := "./gormsqlite.db?_timeout=5000"
	open, err := gorm.Open(sqlite.Open(dbpath))
	if err != nil {
		panic(err)
	}

	SqliteHandler = open
}

func CleanSqlite() {
	os.Remove("./gormsqlite.db")
}
