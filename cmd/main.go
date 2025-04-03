package main

import (
	"github.com/MicroMolekula/auth-service/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth service",
		})
	})

	if err := engine.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
