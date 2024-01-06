package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TokenType string

const (
	RefreshTokenType TokenType = "refresh"
	AccessTokenType  TokenType = "access"
)

type Login struct {
	Phone string `json:"phone" binding:"required"`

	Code        string    `json:"-"`
	Id          int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	IsStaff     bool      `json:"-"`
	IsSuperuser bool      `json:"-"`
	IsActive    bool      `json:"-"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Phone, validation.Required),
	)
}

type LoginResponse struct {
	ID int `json:"id"`
}

type LoginConfirm struct {
	ID   int    `json:"id" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func (l LoginConfirm) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.ID, validation.Required),
		validation.Field(&l.Code, validation.Required),
	)
}

type LoginConfirmResponse struct {
	*User        `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId    int       `json:"user_id"`
	TokenType TokenType `json:"type"`
}
