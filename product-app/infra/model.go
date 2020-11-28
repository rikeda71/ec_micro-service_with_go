package infra

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductId    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	ProductPrice int    `json:"product_price"`
}
