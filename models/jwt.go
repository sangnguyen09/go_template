package models

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	UserId int `json:"userId"`
	Username  string `json:"username"`
	Role Role `json:"role"`
	jwt.StandardClaims
}