package middleware

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type UserRefreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
