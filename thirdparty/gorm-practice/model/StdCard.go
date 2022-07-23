package model

type StuCard struct {
	ID        uint   `gorm:"primary_key"`
	Number    string `gorm:"not null"`
	StudentId uint   `gorm:"not null"`
}
