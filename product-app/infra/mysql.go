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
	PROTOCOL := "tcp(product-db:3306)"
	DBNAME := "micro_product"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	Migrator = DB.Migrator()
	if !Migrator.HasTable(Product{}) {
		_ = Migrator.CreateTable(Product{})
		insertMock()
	}
	_ = DB.AutoMigrate(&Product{})
}

func insertMock() {
	product1 := Product{
		ProductId: 1,
		ProductName: "iPhone",
		ProductImage: "https://mock.com/iphone.png",
		ProductPrice: 50000,
	}
	product2 := Product{
		ProductId:    2,
		ProductName:  "MacBook",
		ProductImage: "https://mock.com/macbook.png",
		ProductPrice: 150000,
	}
	product3 := Product{
		ProductId:    3,
		ProductName:  "IPad",
		ProductImage: "https://mock.com/ipad.png",
		ProductPrice: 100000,
	}
	products := []Product{product1, product2, product3}
	for _, p := range products {
		DB.Create(&p)
		DB.Save(&p)
	}
}