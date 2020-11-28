package infra

import "gorm.io/gorm"

// User ユーザ情報
type User struct {
	gorm.Model
	Email    string `json:"user_id"`
	Password string `json:"password"`
}
