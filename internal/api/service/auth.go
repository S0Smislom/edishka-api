package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"food/internal/api/model"
	"food/internal/api/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	accessTokenTTL int
	tokenSecret    string
	repo           repository.Auth
	userRepo       repository.User
}

func NewAuthService(accessTokenTTL int, tokenSecret string, repo repository.Auth, userRepo repository.User) *AuthService {
	return &AuthService{
		repo:           repo,
		userRepo:       userRepo,
		accessTokenTTL: accessTokenTTL,
		tokenSecret:    tokenSecret,
	}
}

func (s *AuthService) CreateUser(data *model.Login) (*model.LoginResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	generatedCode := generateConfirmationCode()
	data.Code = &generatedCode
	user_id, err := s.repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	return user_id, nil
}

func (s *AuthService) Login(data *model.LoginConfirm) (*model.LoginConfirmResponse, error) {
	// if err := data.Validate(); err != nil {
	// 	return nil, err
	// }
	if err := data.Validate(); err != nil {
		return nil, err
	}
	dbUser, err := s.userRepo.GetById(data.ID)
	if err != nil {
		return nil, err
	}
	if dbUser.Code == nil || generateHash(data.Code) != *dbUser.Code {
		return nil, errors.New("Invalid code")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.accessTokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: dbUser.ID,
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

	response := &model.LoginConfirmResponse{
		AccessToken: access_token,
		User:        dbUser,
	}
	return response, nil
}

func generateConfirmationCode() string {
	code := "1111"
	return generateHash(code)
}

func generateHash(s string) string {
	salt := "hjqrhjqw124617ajfhajs"
	hash := sha1.New()
	hash.Write([]byte(s))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
