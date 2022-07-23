package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID      string `gorm:"not null" gorm:"uuid" faker:"uuid_digit,unique"`
	Name      string `gorm:"not null" faker:"username,unique"`
	Email     string `gorm:"not null" faker:"email,unique"`
	Cellphone string `gorm:"not null" faker:"phone_number"`
	Gender    string `gorm:"not null" json:"gender" faker:"gender"`
	Password  string `gorm:"not null" json:"password" faker:"password"`
	Address   string `gorm:"not null" json:"address" faker:"mac_address"`
	Avatar    string `gorm:"not null" faker:"url"`
}

//CREATE TABLE `users` (`id` int unsigned AUTO_INCREMENT,`name` varchar(255) NOT NULL,`age` int NOT NULL , PRIMARY KEY (`id`))

//ALTER TABLE users ADD INDEX index_name_age (name, age);
//ALTER TABLE users ADD INDEX index_age (age);
//ALTER TABLE users ADD INDEX index_name (name);
type UserInfo struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"not null;default:''"`
	Age        int    `gorm:"not null;default:0"`
	CompanyID  uint   `gorm:"not null"`
	Company    Company
	CreditCard CreditCard `gorm:"foreignKey:UserID"` // has one, 多个则是 has many

	Languages []Language `gorm:"many2many:user_languages;"`
	Orders    []Order    `gorm:"foreignKey:UserID"`
}

func (u *UserInfo) TableName() string {
	return "user"
}

type CreditCard struct {
	ID     uint `gorm:"primarykey"`
	Number string
	UserID uint
}

func (*CreditCard) TableName() string {
	return "creditcard"
}

type Language struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
