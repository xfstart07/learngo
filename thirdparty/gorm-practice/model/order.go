package model

type Order struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Price     float64
	OrderItem OrderItem
}

func (u *Order) TableName() string {
	return "order"
}

type OrderItem struct {
	ID        uint `gorm:"primary_key"`
	ItemName  string
	OrderID   uint
	ProductID uint
	Product   Product
}

func (u *OrderItem) TableName() string {
	return "orderItem"
}

type Product struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

func (u *Product) TableName() string {
	return "product"
}
