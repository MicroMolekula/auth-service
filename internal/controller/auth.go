package controller

import (
	"github.com/MicroMolekula/auth-service/internal/dto"
	"github.com/MicroMolekula/auth-service/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) LoginController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var credential dto.UserLogin
	if err := ctx.ShouldBindBodyWithJSON(&credential); err != nil {
		ErrorResponse(http.StatusBadRequest, "Не правильный формат запроса", err, ctx)
		return
	}
	token, err := ac.authService.Login(credential)
	if err != nil {
		ErrorResponse(http.StatusUnauthorized, "Неверный логин или пароль", err, ctx)
		return
	}
	session.Set("SESSID", token.RefreshToken)
	if err = session.Save(); err != nil {
		ErrorResponse(http.StatusInternalServerError, "Ошибка сервера", err, ctx)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
}

func (ac *AuthController) RegisterController(ctx *gin.Context) {
	var credential dto.UserRegister
	if err := ctx.ShouldBindBodyWithJSON(&credential); err != nil {
		ErrorResponse(http.StatusBadRequest, "Не правильный формат запроса", err, ctx)
		return
	}
	token, err := ac.authService.Register(credential)
	if err != nil {
		ErrorResponse(http.StatusConflict, "Такой пользователь уже зарегистрирован", err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token":        token.Token,
		"refreshToken": token.RefreshToken,
	})
}
