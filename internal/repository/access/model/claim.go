package model

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	ExamplePath = "/auth_v1.AuthV1/Get"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserName string `json:"username"`
	Role     string `json:"role"`
}
