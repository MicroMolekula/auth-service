package service

import (
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/repository"
)

type AuthService struct {
	jwtService     *JWTService
	userRepository *repository.UserRepository
	cfg            *config.Config
}
