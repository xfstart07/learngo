// Author: xufei
// Date: 2020/10/26

package main

import (
	"fmt"
	"sync"
	"time"

	db2 "learngo/pkg/gormexp/db"
	"learngo/pkg/gormexp/model"
)

// 经实验，在 MySQL 5.7 版本下，隔离级别：可重复读，并没有发送幻读，一个事务内两次读取的数据一致。
// 为什么？
// 在 MySQl 5.7 版本的 RR 级别下，幻读已被解决。加入 Next-key lock 间隙锁。
// https://www.cnblogs.com/michael9/p/12358631.html

func main() {
	db2.InitMySql()

	db := db2.MySqlHandler
	var sy sync.WaitGroup

	var len1, len2 int
	sy.Add(1)
	go func() {
		defer sy.Done()

		tx := db.Begin()

		var stus []model.Student
		tx.Find(&stus)
		len1 = len(stus)
		fmt.Println(len1)

		time.Sleep(100 * time.Microsecond)

		// 出现幻读时，第二次读到的值和第一次不一样
		var stus2 []model.Student
		tx.Find(&stus2)
		len2 = len(stus2)
		fmt.Println(len2)

		tx.Commit()
	}()

	sy.Add(1)
	go func() {
		defer sy.Done()

		tx := db.Begin()

		stu := model.Student{
			Name:    "学生Read",
			ClassId: 5,
		}
		err := tx.Create(&stu).Error
		if err != nil {
			fmt.Println("错误", err)
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	sy.Wait()

	if len1 != len2 {
		fmt.Println("出现幻读！")
	}
}
