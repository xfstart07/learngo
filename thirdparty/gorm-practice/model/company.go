package model

type Company struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

func (u *Company) TableName() string {
	return "company"
}
