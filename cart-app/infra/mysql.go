package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Migrator gorm.Migrator

// GormConnect gormによってmysqlに接続するためのメソッド
func GormConnect() {
	USER := "root"
	PASS := "mysql"
	PROTOCOL := "tcp(cart-db:3306)"
	DBNAME := "micro_cart"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	Migrator = DB.Migrator()
	if !Migrator.HasTable(Cart{}) {
		_ = Migrator.CreateTable(Cart{})
	}
	_ = DB.AutoMigrate(&Cart{})
}

