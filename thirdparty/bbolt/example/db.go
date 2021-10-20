package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.etcd.io/bbolt"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Created int64  `json:"created"`
}

func (User) TableName() string {
	return "user"
}

func OpenDB(path string) *bbolt.DB {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Create(db *bbolt.DB) {
	db.Update(func(tx *bbolt.Tx) error {
		var user User
		b, _ := tx.CreateBucketIfNotExists([]byte(user.TableName()))
		for i := 0; i < 100; i++ {
			user = User{
				Name: fmt.Sprintf("name%d", i),
				Age:  rand.Intn(100),
			}
			id, _ := b.NextSequence()
			user.Id = int(id)
			body, _ := json.Marshal(user)

			err := b.Put(itob(user.Id), body)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func Get(db *bbolt.DB) {
	db.View(func(tx *bbolt.Tx) error {
		user := User{}
		b := tx.Bucket([]byte(user.TableName()))

		v := b.Get(itob(2))
		fmt.Printf("%v\n", string(v))
		return nil
	})
}

func FindRange(db *bbolt.DB) {
	db.View(func(tx *bbolt.Tx) error {
		user := User{}
		b := tx.Bucket([]byte(user.TableName())).Cursor()

		// 查找key在2到5范围的数据
		min := itob(2)
		max := itob(5)
		fmt.Println("min:", string(min), "max", string(max))

		for k, v := b.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = b.Next() {
			fmt.Printf("k=%s, v=%s\n", string(k), string(v))
		}
		return nil
	})
}

func CreateByTime(db *bbolt.DB) {
	db.Update(func(tx *bbolt.Tx) error {
		var user User
		b, _ := tx.CreateBucketIfNotExists([]byte(user.TableName()))
		for i := 1; i < 100; i++ {
			user = User{
				Id:      i,
				Name:    fmt.Sprintf("name%d", i),
				Age:     rand.Intn(100),
				Created: time.Now().Unix(),
			}
			// 用时间戳+ID作为key，1634728694:1
			key := fmt.Sprintf("%v:%v", user.Created, user.Id)
			body, _ := json.Marshal(user)

			err := b.Put([]byte(key), body)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func FindRangeByTime(db *bbolt.DB) {
	db.View(func(tx *bbolt.Tx) error {
		user := User{}
		b := tx.Bucket([]byte(user.TableName())).Cursor()

		// 时间戳查询
		min := []byte(fmt.Sprintf("%v", time.Now().Add(-1*time.Minute).Unix()))
		max := []byte(fmt.Sprintf("%v", time.Now().Add(time.Minute).Unix()))
		fmt.Println("min:", string(min), "max", string(max))

		for k, v := b.Seek(min); k != nil; k, v = b.Next() {
			fmt.Printf("%s, %s\n", string(k), string(v))
		}
		return nil
	})
}

func FindEach(db *bbolt.DB) {
	db.View(func(tx *bbolt.Tx) error {
		user := User{}
		b := tx.Bucket([]byte(user.TableName()))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", string(k), string(v))
			return nil
		})
		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
