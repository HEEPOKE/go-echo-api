package models

import "github.com/golang-jwt/jwt"

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
