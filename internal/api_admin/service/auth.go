package service

import (
	"errors"
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
	"food/pkg/hash"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt string = "ljkfaslgratljhsdflkjgzxlawlsur"
)

type AuthService struct {
	tokenTTL    int
	tokenSecret string
	userRepo    repository.User
}

func NewAuthService(tokenTTL int, tokenSecret string, userRepo repository.User) *AuthService {
	return &AuthService{
		tokenTTL:    tokenTTL,
		tokenSecret: tokenSecret,
		userRepo:    userRepo,
	}
}

func (s *AuthService) Create(data *model.CreateUser) (int, error) {
	if err := data.Validate(); err != nil {
		return 0, err
	}
	if data.Password != data.Password2 {
		return 0, errors.New("Пароли не совпали")
	}
	data.Password = hash.GenerateHash(data.Password, salt)
	return s.userRepo.Create(data)
}

func (s *AuthService) Login(data *model.Login) (*model.LoginResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByPhone(data.Login)
	if err != nil {
		return nil, err
	}
	if user.Password == nil || *user.Password != hash.GenerateHash(data.Password, salt) {
		return nil, errors.New("Wrong password")
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.tokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	access_token, err := accessToken.SignedString([]byte(s.tokenSecret))
	if err != nil {
		return nil, err
	}

	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(s.re),
	// 	},
	// })

	response := &model.LoginResponse{
		AccessToken: access_token,
		User:        user,
	}
	return response, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.tokenSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, nil
}
