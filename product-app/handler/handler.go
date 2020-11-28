package handler

import (
	"github.com/labstack/echo"
	"github.com/s14t284/ec_micro-service_with_go/infra"
)

// ErrorResponse is
type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"error_message"`
}

// IndexHandler is
func IndexHandler(c echo.Context) error {
	return c.String(200, "Welcome to Catalog Service!")
}

// GetAllProductsHandler 商品を全取得
func GetAllProductsHandler(c echo.Context) error {
	var products []infra.Product
	infra.DB.Find(&products)
	if len(products) > 0 {
		return c.JSON(200, products)
	}
	return c.JSON(400, ErrorResponse{1, "products don't exists"})
}
