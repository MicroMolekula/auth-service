package main

import (
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/MicroMolekula/auth-service/internal/controller"
	"github.com/MicroMolekula/auth-service/internal/database"
	"github.com/MicroMolekula/auth-service/internal/repository"
	"github.com/MicroMolekula/auth-service/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(db)
	jwtService := service.NewJwtService(cfg)
	authService := service.NewAuthService(jwtService, userRepository, cfg)
	authController := controller.NewAuthController(authService)
	engine := gin.Default()
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{
		HttpOnly: true,
	})
	engine.Use(sessions.Sessions("FITSESSION", store))
	engine.POST("/login", authController.Login)
	engine.POST("/register", authController.Register)
	engine.GET("/refresh_token", authController.RefreshToken)
	engine.GET("/logout", authController.Logout)

	if err := engine.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
