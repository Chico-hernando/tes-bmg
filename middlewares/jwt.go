package middlewares

import (
	"bmg/constants"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.JWT_SECRET))
}

func GetUserIdFromJWT(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}