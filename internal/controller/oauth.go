package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/dto"
	"github.com/MicroMolekula/auth-service/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

type OauthYandexController struct {
	cfg         *config.Config
	oauthConfig *oauth2.Config
	authService *service.AuthService
}

func NewOauthYandexController(cfg *config.Config, oauthConfig *oauth2.Config, authService *service.AuthService) *OauthYandexController {
	return &OauthYandexController{
		cfg:         cfg,
		oauthConfig: oauthConfig,
		authService: authService,
	}
}

func (oc *OauthYandexController) LoginHandler(ctx *gin.Context) {
	url := oc.oauthConfig.AuthCodeURL(oc.cfg.OauthYandex.State)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (oc *OauthYandexController) CallbackHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	state := ctx.Query("state")
	if state != oc.cfg.OauthYandex.State {
		ErrorResponse(http.StatusBadRequest, "Invalid state", errors.New("invalid state"), ctx)
		return
	}
	code := ctx.Query("code")
	if code == "" {
		ErrorResponse(http.StatusBadRequest, "Invalid code", errors.New("invalid code"), ctx)
		return
	}
	token, err := oc.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, "Failed to exchange token", errors.New("failed to exchange token"), ctx)
		return
	}
	client := oc.oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://login.yandex.ru/info?format=json")
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, "Failed to get user info", errors.New("failed to get user info"), ctx)
		return
	}
	defer resp.Body.Close()

	var yandexUser dto.YandexUser
	if err := json.NewDecoder(resp.Body).Decode(&yandexUser); err != nil {
		ErrorResponse(http.StatusInternalServerError, "Failed to parse user info", errors.New("failed to parse user info"), ctx)
		return
	}

	tokenDto, err := oc.authService.LoginWithYandexId(&yandexUser)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, "Failed to get user info", errors.New("failed to get user info"), ctx)
		return
	}
	session.Set("FITSESSION", tokenDto.RefreshToken)
	session.Save()
	ctx.Redirect(http.StatusTemporaryRedirect, "https://not-five.ru")
}
