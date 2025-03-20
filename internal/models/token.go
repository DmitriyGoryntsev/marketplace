package models

import "github.com/golang-jwt/jwt"

type AccessTokenClaims struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
