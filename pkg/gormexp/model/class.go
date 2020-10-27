// Author: xufei
// Date: 2020/10/20

package model

type Class struct {
	ID       uint      `gorm:"primary_key"`
	Name     string    `gorm:"not null" json:"name"`
	Students []Student `gorm:"foreignKey:ClassId"` // has many
}
