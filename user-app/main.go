package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email string `json:"user_id"`
	Password string `json:"password"`
}

func gormConnect() *gorm.Migrator {
	USER := "root"
	PASS := "mysql"
	PROTOCOL := "tcp(user-db:3306)"
	DBNAME := "micro_user"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrator := db.Migrator()
	if !migrator.HasTable(User{}) {
		migrator.CreateTable(User{})
	}

	return &migrator
}

func main() {
	gormConnect()
}
