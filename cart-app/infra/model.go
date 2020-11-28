package infra

import "gorm.io/gorm"

// Cart カート情報
type Cart struct {
	gorm.Model
	UserId    int `json:"user_id"`
	ProductId int `json:"product_id"`
}
