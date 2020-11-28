package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/s14t284/ec_micro-service_with_go/infra"
	"net/http"
	"strconv"
)

const authHeader = "Authorization"

// User ユーザ情報
type User struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
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
	return c.String(200, "Welcome to Cart Service!")
}

// GetALlCartItemHandler カート情報を全取得
func GetAllCartItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}

	var carts []infra.Cart
	infra.DB.Where("user_id = ?", user.UserId).Find(&carts)
	if len(carts) > 0 {
		return c.JSON(200, carts)
	}
	return c.JSON(200, []infra.Cart{})
}

// CrateCartItemHandler カートに商品を追加
func CreateCartItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}

	productId, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return c.JSON(400, ErrorResponse{2, fmt.Sprintf("product id format is invalid: [error][%s]", err)})
	}
	cart := infra.Cart{
		UserId:    user.UserId,
		ProductId: productId,
	}
	infra.DB.Create(&cart)
	return c.JSON(200, cart)
}

// DeleteCartItemHandler カート商品削除
func DeleteCartItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}
	cartId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, ErrorResponse{2, fmt.Sprintf("cart id format is invalid: [error][%s]", err)})
	}
	var cart infra.Cart
	cart.ID = uint(cartId)
	cart.UserId = user.UserId
	infra.DB.First(&cart)
	infra.DB.Delete(&cart)
	return c.JSON(200, SuccessfulResponse{Message: fmt.Sprintf("delete cart item %d is success", cartId)})
}

// DeleteCartItemHandler カート商品全削除
func DeleteAllCartItemHandler(c echo.Context) error {
	// validate token
	user, err := validateJwtToken(c.Request())
	if err != nil {
		return c.JSON(400, ErrorResponse{
			ErrorCode: 1,
			Message:   fmt.Sprintf("require validate token [error][%s]", err),
		})
	}
	var carts []infra.Cart
	infra.DB.Where("user_id = ?", user.UserId).Find(&carts)
	infra.DB.Delete(&carts)
	return c.JSON(200, SuccessfulResponse{
		Message: fmt.Sprintf("delete all cart of '%s' user item is success [user_id][%d]", user.Email, user.UserId),
	})
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
