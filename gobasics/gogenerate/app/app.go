package app

import "fmt"

//go:generate mockgen -destination=../mock/app.go -package=mock -source=app.go App

type App interface {
	Start() error
	Apply() error
}

type app struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *app) Apply() error {
	fmt.Println(u)
	return nil
}

func (u *app) Start() error {
	u.Name = "edgeComputing"
	u.Age = 10
	return nil
}
