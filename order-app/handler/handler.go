package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/s14t284/ec_micro-service_with_go/infra"
	"net/http"
)

const authHeader = "Authorization"

// User ユーザ情報
type User struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

// OrderDetailArray 注文詳細情報
type OrderDetailArray struct {
	OrderDetails []infra.OrderDetail `json:"order_details" binding:"required"`
}

// SuccessfulResponse is
type SuccessfulResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is
type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"error_message"`
}

// IndexHandler is
func IndexHandler(c echo.Context) error {
	return c.String(200, "Welcome to Order Service!")
}

// CreateOrderItemHandler 注文を追加する
func CreateOrderItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}

	details := new(OrderDetailArray)
	if err = c.Bind(details); err != nil {
		return c.JSON(400, ErrorResponse{1, fmt.Sprintf("failed to parse order information [error][%s]", err)})
	}

	fmt.Printf("result of parse order: %#v\n", details.OrderDetails)
	order := infra.Order{UserId: user.UserId, OrderDetails: details.OrderDetails}
	infra.DB.Create(&order)

	return c.JSON(200, order)
}

// GetAllOrderItemHandler 注文全取得
func GetAllOrderItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}

	var orders []infra.Order
	infra.DB.Where("user_id = ?", user.UserId).Find(&orders)
	for i, o := range orders {
		infra.DB.Model(&o).Association("OrderDetails").Find(&orders[i].OrderDetails)
	}
	if len(orders) > 0 {
		return c.JSON(200, orders)
	}

	return c.JSON(400, ErrorResponse{Message: "order items are not found"})
}

func validateJwtToken(req *http.Request) (*User, error) {
	fmt.Println("validate Request Token:" + req.Header.Get(authHeader))
	client := resty.New()
	resp, err := client.R().SetHeader(authHeader, req.Header.Get(authHeader)).Get("http://user-app:3000/me")
	if err != nil || resp.StatusCode() != 200 {
		fmt.Printf("failed to signin with jwt token: %s\n", err)
		return nil, err
	}

	user := new(User)
	fmt.Println("validation result: " + resp.String())
	if err := json.Unmarshal([]byte(resp.String()), user); err != nil {
		fmt.Printf("json parse error: %s\n", err)
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}
