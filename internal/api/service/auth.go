package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
	"food/pkg/hash"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt string = "hjqrhjqw124617ajfhajs"
)

type AuthService struct {
	accessTokenTTL  int
	refreshTokenTTL int
	tokenSecret     string
	repo            repository.Auth
	userRepo        repository.User
}

func NewAuthService(accessTokenTTL, refreshTokenTTL int, tokenSecret string, repo repository.Auth, userRepo repository.User) *AuthService {
	return &AuthService{
		repo:            repo,
		userRepo:        userRepo,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		tokenSecret:     tokenSecret,
	}
}

func (s *AuthService) CreateUser(data *model.Login) (*model.LoginResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	data.Code = generateConfirmationCode()

	userId, err := s.userRepo.GetByPhone(data.Phone)
	if err != nil {
		userId, err = s.repo.CreateUser(data)
		if err != nil {
			return nil, err
		}
	} else {
		if err := s.userRepo.UpdateCode(userId, data.Code); err != nil {
			return nil, err
		}
	}
	return &model.LoginResponse{ID: userId}, nil
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
		return nil, &exceptions.UnauthorizedError{Msg: "Invalid code"}
	}
	// Generate access token
	accessTokenSigned, err := generateToken(model.AccessTokenType, dbUser, s.accessTokenTTL, s.tokenSecret)
	if err != nil {
		return nil, err
	}
	// Generate refresh token
	refreshTokenSigned, err := generateToken(model.RefreshTokenType, dbUser, s.refreshTokenTTL, s.tokenSecret)
	if err != nil {
		return nil, err
	}

	response := &model.LoginConfirmResponse{
		AccessToken:  accessTokenSigned,
		RefreshToken: refreshTokenSigned,
		User:         dbUser,
	}
	return response, nil
}

func (s *AuthService) ParseToken(accessToken string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &exceptions.UnauthorizedError{Msg: "invalid signing method"}
		}

		return []byte(s.tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return nil, &exceptions.UnauthorizedError{Msg: "token claims are not of type *TokenClaims"}
	}

	return claims, nil
}
func (s *AuthService) Refresh(userId int) (*model.LoginConfirmResponse, error) {
	dbUser, err := s.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}
	// Generate access token
	accessTokenSigned, err := generateToken(model.AccessTokenType, dbUser, s.accessTokenTTL, s.tokenSecret)
	if err != nil {
		return nil, err
	}
	// Generate refresh token
	refreshTokenSigned, err := generateToken(model.RefreshTokenType, dbUser, s.refreshTokenTTL, s.tokenSecret)
	if err != nil {
		return nil, err
	}
	response := &model.LoginConfirmResponse{
		AccessToken:  accessTokenSigned,
		RefreshToken: refreshTokenSigned,
		User:         dbUser,
	}
	return response, nil
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

func generateToken(tokenType model.TokenType, dbUser *model.User, ttl int, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: getTokenExpiresAt(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    dbUser.ID,
		TokenType: tokenType,
	})
	tokenSigned, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}

func getTokenExpiresAt(ttl int) time.Time {
	return time.Now().Add(time.Duration(ttl) * time.Hour)
}
