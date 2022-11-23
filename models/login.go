package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Login
type Login struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Claim
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
