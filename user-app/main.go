package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	h "github.com/s14t284/ec_micro-service_with_go/handler"
	"github.com/s14t284/ec_micro-service_with_go/infra"
)

//-------------
// RestAPI
//-------------
func launchRestApi() {
	e := echo.New()
	// CORS
	e.Use(middleware.CORS())
	e.GET("/", h.IndexHandler)
	e.POST("/login", h.LoginHandler)
	e.POST("/user", h.CreateUserHandler)
	e.GET("/me", h.CurrentUserHandler)
	_ = e.Start(":3000")
}

func main() {
	infra.GormConnect()
	launchRestApi()
}
