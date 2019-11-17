package models

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	UserId int `json:"userId"`
	Username  string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}