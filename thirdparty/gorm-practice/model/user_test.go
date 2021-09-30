// Author: xufei
// Date: 2020/7/21

package model

import (
	"testing"
)

//NOTE: save 和 update 效率差不别不大，主要的瓶颈还是 mysql 上
func BenchmarkUpdateBySave(t *testing.B) {
	db := GetDB()

	var users []User
	db.Find(&users)

	lenUser := len(users)
	t.Log(lenUser)
	t.N = 1000
	for i := 0; i < t.N; i++ {
		//k := rand.Intn(lenUser)

		db.Save(&users[i])
	}
	//BenchmarkUpdateBySave-8   	    1000	  13778820 ns/op
	//BenchmarkUpdateBySave-8   	    1000	  10692361 ns/op
	//BenchmarkUpdateBySave-8   	    1000	   8251996 ns/op
}

func BenchmarkUpdate(t *testing.B) {
	db := GetDB()

	var users []User
	db.Find(&users)

	lenUser := len(users)
	t.Log(lenUser)
	t.N = 1000
	for i := 0; i < t.N; i++ {
		//k := rand.Intn(lenUser)

		db.Model(User{}).Where("id=?", users[i].ID).Update("gender", "Male")
	}
	//BenchmarkUpdate-8   	    1000	  10792708 ns/op
	//BenchmarkUpdate-8   	    1000	   9006988 ns/op
}
