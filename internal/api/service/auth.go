package service

import (
	"errors"
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/hash"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt string = "hjqrhjqw124617ajfhajs"
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
	if err := data.Validate(); err != nil {
		return nil, err
	}
	dbUser, err := s.userRepo.GetById(data.ID)
	if err != nil {
		return nil, err
	}
	if dbUser.Code == nil || hash.GenerateHash(data.Code, salt) != *dbUser.Code {
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

func generateConfirmationCode() string {
	code := "1111"
	// return generateHash(code)
	return hash.GenerateHash(code, salt)
}

// func generateHash(s string) string {
// 	salt := "hjqrhjqw124617ajfhajs"
// 	hash := sha1.New()
// 	hash.Write([]byte(s))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }
