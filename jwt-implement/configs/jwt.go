package configs

import (
	"github.com/golang-jwt/jwt/v5"
)


type JwtClaims struct {
	Username string
	jwt.RegisteredClaims
}