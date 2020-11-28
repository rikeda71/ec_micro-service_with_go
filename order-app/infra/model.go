package infra

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId       int           `json:"user_id"`
	OrderDetails []OrderDetail `gorm:"foreignkey:OrderId" ;json:"order_details"`
}

// OrderDetail 注文詳細情報
type OrderDetail struct {
	gorm.Model
	OrderId      int `json:"order_id"`
	ProductId    int `json:"product_id"`
	ProductPrice int `json:"product_price"`
}
