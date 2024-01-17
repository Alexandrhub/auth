package model

import (
	"github.com/golang-jwt/jwt/v5"
)

const ExamplePath = "/user/v1/create"

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
