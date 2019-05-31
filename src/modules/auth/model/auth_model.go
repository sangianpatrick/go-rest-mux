package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Auth contains auth property
type Auth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// BearerClaims contains authorized token
type BearerClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Credential contain name and token
type Credential struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
