package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	timeStandardLayout = "2006-01-02 15:04:05"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) MarshalJSON() (data []byte, err error) {
	type UserAlias User

	userStruct := &struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		*UserAlias
	}{
		UserAlias: (*UserAlias)(u),
	}

	userStruct.CreatedAt = u.CreatedAt.Format(timeStandardLayout)
	userStruct.UpdatedAt = u.UpdatedAt.Format(timeStandardLayout)

	return json.Marshal(userStruct)
}

func (u *User) UnmarshalJSON(data []byte) (err error) {
	type UserAlias User

	userStruct := &struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		*UserAlias
	}{
		UserAlias: (*UserAlias)(u),
	}

	if err = json.Unmarshal(data, &userStruct); err != nil {
		return err
	}

	u.CreatedAt, err = time.Parse(timeStandardLayout, userStruct.CreatedAt)
	if err != nil {
		return err
	}
	u.UpdatedAt, err = time.Parse(timeStandardLayout, userStruct.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	user := &User{
		ID:        1,
		Name:      "Leon",
		Password:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := json.Marshal(&user)
	fmt.Println("编码", string(data), err)

	user2 := &User{}
	err = json.Unmarshal(data, &user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("解码 %+v\n", user2)
}
