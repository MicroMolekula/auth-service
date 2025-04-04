package controller

import (
	"errors"
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

func (ac *AuthController) Login(ctx *gin.Context) {
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
	session.Set("FITSESSION", token.RefreshToken)
	if err = session.Save(); err != nil {
		ErrorResponse(http.StatusInternalServerError, "Ошибка сервера", err, ctx)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
}

func (ac *AuthController) Register(ctx *gin.Context) {
	session := sessions.Default(ctx)
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
	session.Set("FITSESSION", token.RefreshToken)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token.Token,
	})
}

func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	session := sessions.Default(ctx)
	refreshToken := session.Get("FITSESSION")
	if refreshToken == nil {
		ErrorResponse(http.StatusUnauthorized, "Пользователь не автаризован", errors.New("unauthorized"), ctx)
		return
	}
	token, err := ac.authService.RefreshToken(refreshToken.(string))
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, "Ошибка на сервере", err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("FITSESSION")
	if err := session.Save(); err != nil {
		ErrorResponse(http.StatusInternalServerError, "Ошибка на сервере", err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
