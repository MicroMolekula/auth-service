package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/dto"
	"github.com/MicroMolekula/auth-service/internal/service"
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
	//code := ctx.Query("code")
	//if code == "" {
	//	ErrorResponse(http.StatusBadRequest, "Invalid code", errors.New("invalid code"), ctx)
	//	return
	//}
	//token, err := oc.oauthConfig.Exchange(context.Background(), code)
	//if err != nil {
	//	ErrorResponse(http.StatusInternalServerError, "Failed to exchange token", errors.New("failed to exchange token"), ctx)
	//	return
	//}
	token := ctx.Query("access_token")
	if token == "" {
		ErrorResponse(http.StatusBadRequest, "Invalid access token", errors.New("invalid access token"), ctx)
		return
	}
	accessToken, err := oc.oauthConfig.Exchange(context.Background(), token)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, "Failed to exchange token", errors.New("failed to exchange token"), ctx)
		return
	}
	client := oc.oauthConfig.Client(context.Background(), accessToken)
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
	fmt.Printf("%+v\n", yandexUser)
	ctx.Redirect(http.StatusTemporaryRedirect, "https://not-five.ru")
}
