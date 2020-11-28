package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

//-------------
// RestAPI
//-------------
func launchRestApi() {
	e := echo.New()
	// CORS
	e.Use(middleware.CORS())
	e.GET("/", IndexHandler)
	_ = e.Start(":3000")
}

func IndexHandler(c echo.Context) error {
	return c.String(200, "Welcome to User Service!")
}

type ErrorResponse struct {
	ErrorCode int `json:"error_code"`
	Message string `json:"error_message"`
}

func main() {
	gormConnect()
	launchRestApi()
}
