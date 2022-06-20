package model

type Company struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
