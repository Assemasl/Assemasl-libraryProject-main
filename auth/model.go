package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthorID int    `json:"author_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUser struct {
	ID       int
	Username string
	Password string
	AuthorID int
}

type Claims struct {
	Username string `json:"username"`
	AuthorID int    `json:"author_id"`
	jwt.RegisteredClaims
}
