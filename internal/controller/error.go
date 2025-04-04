package controller

import "github.com/gin-gonic/gin"

func ErrorResponse(status int, message string, err error, ctx *gin.Context) {
	ctx.JSON(status, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}
