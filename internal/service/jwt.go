package service

import (
	"errors"
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/dto"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrorInvalidJWTToken = errors.New("invalid JWT token")
)

type DataJwt struct {
	Id    int    `json:"id"`
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
	Role  string `json:"role"`
	Email string `json:"username"`
	jwt.RegisteredClaims
}

type JWTService struct {
	cfg *config.Config
}

func NewJwtService(cfg *config.Config) *JWTService {
	return &JWTService{cfg: cfg}
}

func (s *JWTService) GenerateTokenByUser(user *dto.User, ttl int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, DataJwt{
		Exp:   time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
		Iat:   time.Now().Unix(),
		Id:    user.Id,
		Email: user.Email,
	})

	token, err := jwtToken.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *JWTService) CreateTokenByRefreshToken(refreshToken string, ttl int) (string, error) {
	dataRefreshToken, err := s.ParseToken(refreshToken)
	if err != nil {
		return "", err
	}
	user := &dto.User{
		Id:    dataRefreshToken.Id,
		Email: dataRefreshToken.Email,
	}
	return s.GenerateTokenByUser(user, ttl)
}

func (s *JWTService) ParseToken(tokenString string) (*DataJwt, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DataJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	result, ok := token.Claims.(*DataJwt)
	if !ok {
		return nil, ErrorInvalidJWTToken
	}

	return result, nil
}
