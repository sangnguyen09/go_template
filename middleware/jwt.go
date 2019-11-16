package middleware

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sangnguyen09/go_template/config"
	"github.com/sangnguyen09/go_template/models"
)

func JWTMiddleware() echo.MiddlewareFunc {
	configs := middleware.JWTConfig{
		Claims:     &models.JWTCustomClaims{},
		SigningKey: []byte(fmt.Sprintf("%s", config.Config.Encryption.JWTSecret)),
	}
	return middleware.JWTWithConfig(configs)
}

func GenToken(user models.User) (string, error) {
	if user.UserId == 0 || len(user.Role) == 0 || len(user.Username) == 0 {
		return "", errors.New("UserId == 0 || len(Role) == 0 || len(Phone) == 0")
	}

	timeExp := time.Duration(config.Config.Encryption.JWTExp)

	claims := &models.JWTCustomClaims{
		UserId:   user.UserId,
		Role:     user.Role,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * timeExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(fmt.Sprintf("%s", config.Config.Encryption.JWTSecret)))
	if err != nil {
		return "", err
	}

	return t, nil
}
func GenTokenRefresh(user models.User) (string, error) {
	if user.UserId == 0 || len(user.Role) == 0 || len(user.Username) == 0 {
		return "", errors.New("UserId == 0 || len(Role) == 0 || len(Phone) == 0")
	}
	timeExp := time.Duration(config.Config.Encryption.JWTExp)
	claims := &models.JWTCustomClaims{
		UserId:   user.UserId,
		Role:     user.Role,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * timeExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(fmt.Sprintf("%s", config.Config.Encryption.JWTExpRefresh)))
	if err != nil {
		return "", err
	}

	return t, nil
}
func ParseJWTToken(tokenString string) *models.JWTCustomClaims {
	mySigningKey := []byte(fmt.Sprintf("%s", config.Config.Encryption.JWTExpRefresh))

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTCustomClaims{}, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(*models.JWTCustomClaims); ok {
		return claims
	}

	return nil
}
