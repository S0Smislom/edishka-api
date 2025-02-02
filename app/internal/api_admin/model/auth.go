package model

import (
	"regexp"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Login, validation.Required, validation.Match(regexp.MustCompile("^(\\+)[0-9]{11}$"))),
		validation.Field(&l.Password, validation.Required),
	)
}

type LoginResponse struct {
	*User        `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
