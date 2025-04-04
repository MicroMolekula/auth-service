package service

import (
	"errors"
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/dto"
	"github.com/MicroMolekula/auth-service/internal/repository"
	"github.com/MicroMolekula/auth-service/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	jwtService     *JWTService
	userRepository *repository.UserRepository
	cfg            *config.Config
}

func NewAuthService(jwtService *JWTService, userRepository *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		jwtService:     jwtService,
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (s *AuthService) Login(creds dto.UserLogin) (*dto.Token, error) {
	user, err := s.userRepository.FindOneByEmail(creds.Email)
	if err != nil {
		return nil, err
	}
	if err := utils.CheckPassword(user.Password, creds.Password); err != nil {
		return nil, err
	}
	userDto := dto.UserModelToDto(user)
	token, err := s.jwtService.GenerateTokenByUser(userDto, s.cfg.JWT.TTL)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwtService.GenerateTokenByUser(userDto, s.cfg.JWT.RefreshTTL)
	if err != nil {
		return nil, err
	}
	return &dto.Token{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Register(creds dto.UserRegister) (*dto.Token, error) {
	user, err := s.userRepository.FindOneByEmail(creds.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("email already registered")
	}
	password, err := utils.HashPassword(creds.Password)
	if err != nil {
		return nil, err
	}
	userModel := dto.UserRegisterToModel(&dto.UserRegister{
		Name:     creds.Name,
		Email:    creds.Email,
		Password: password,
	})
	if err = s.userRepository.Create(userModel); err != nil {
		return nil, err
	}
	return s.Login(dto.UserLogin{
		Email:    creds.Email,
		Password: creds.Password,
	})
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	token, err := s.jwtService.CreateTokenByRefreshToken(refreshToken, s.cfg.JWT.TTL)
	if err != nil {
		return "", err
	}
	return token, nil
}
