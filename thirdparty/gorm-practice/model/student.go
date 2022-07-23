package model

type Student struct {
	ID      uint    `gorm:"primary_key"`
	Name    string  `json:"name"`
	ClassId uint    `gorm:"not null" json:"class_id"`
	Class   Class   `gorm:"foreignKey:ClassId"`   // belongs to
	StuCard StuCard `gorm:"foreignKey:StudentId"` // has one
	Grade   int     `gorm:"not null;default:1"`
}
