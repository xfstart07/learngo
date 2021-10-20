package main

import (
	"os"
	"testing"
)

const testPath = "kv.db"

func TestFindEach(t *testing.T) {
	db := OpenDB(testPath)
	defer os.Remove(testPath)

	Create(db)

	FindEach(db)
}

func TestFindRange(t *testing.T) {
	db := OpenDB(testPath)
	defer os.Remove(testPath)

	Create(db)

	FindRange(db)
}

func TestFindRangeByTime(t *testing.T) {
	db := OpenDB(testPath)
	defer os.Remove(testPath)

	CreateByTime(db)
	FindRangeByTime(db)
}
