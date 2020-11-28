package handler

import (
	"github.com/labstack/echo"
)

// ErrorResponse is
type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"error_message"`
}

// IndexHandler is
func IndexHandler(c echo.Context) error {
	return c.String(200, "Welcome to Order Service!")
}
