package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
