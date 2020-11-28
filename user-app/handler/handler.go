package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"github.com/s14t284/ec_micro-service_with_go/infra"
	"time"
)

const (
	secretKey = "hoge"
)

// SuccessfulResponse is
type SuccessfulResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is
type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"error_message"`
}

// SignedResponse jwt認証に成功した時に返却するレスポンス
type SignedResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

// JwtToken jwtトークンを格納するオブジェクト
type JwtToken struct {
	Token string `json:"token"`
}

// IndexHandler is
func IndexHandler(c echo.Context) error {
	return c.String(200, "Welcome to User Service!")
}

// LoginHandler ログイン処理を行うハンドラー
func LoginHandler(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")
	// ログイン
	var user infra.User
	if infra.DB.Where(infra.User{Email: email, Password: pass}).First(&user).Error == nil {
		// make JWT
		token, err := generateJwtToken(user)
		if err != nil {
			return c.JSON(400, ErrorResponse{1, fmt.Sprintf("failed to generate jwt: %s", err.Error())})
		}
		return c.JSON(200, JwtToken{Token: token})
	} else {
		return c.JSON(400, ErrorResponse{1, "failed to login"})
	}
}

// CreateUserHandler ユーザ登録を行うハンドラー
func CreateUserHandler(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")
	// ログイン
	user := infra.User{
		Email:    email,
		Password: pass,
	}
	infra.DB.Create(&user)
	return c.JSON(200, SuccessfulResponse{Message: "create user is success"})
}

// CurrenctUserHandler ユーザ情報を取得するためのハンドラー
func CurrentUserHandler(c echo.Context) error {
	// 署名の検証
	token, err := request.ParseFromRequest(c.Request(), request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})
	if err != nil {
		_ = c.JSON(400, ErrorResponse{1, fmt.Sprintf("failed to signed token: %s", err)})
	}
	// claims := token.Claims.(jwt.MapClaims)
	claims := token.Claims.(jwt.MapClaims)
	res := SignedResponse{
		UserId: int(claims["user_id"].(float64)),
		Email: claims["email"].(string),
	}
	return c.JSON(200, res)
}

func generateJwtToken(user infra.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// make token
	token.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	// add signature
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt token [error][%s]", err)
	}
	return tokenString, nil
}
