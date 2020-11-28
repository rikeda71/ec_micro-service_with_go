package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	h "github.com/s14t284/ec_micro-service_with_go/handler"
	"github.com/s14t284/ec_micro-service_with_go/infra"
)

func launchRestApi() {
	e := echo.New()
	// CORS
	e.Use(middleware.CORS())
	e.GET("/", h.IndexHandler)
	e.GET("/cart", h.GetAllCartItemHandler)
	e.POST("/cart", h.CreateCartItemHandler)
	e.DELETE("/cart/:id", h.DeleteCartItemHandler)
	e.DELETE("/cart", h.DeleteAllCartItemHandler)
	_ = e.Start(":3000")
}

func main() {
	infra.GormConnect()
	launchRestApi()
}
