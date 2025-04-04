package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GetTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
	if !ok {
		authHeader = ctx.Request.Header.Get("Sec-WebSocket-Protocol")
		tokenString, ok = strings.CutPrefix(authHeader, "auth, ")
		if !ok {
			return "", errors.New("invalid token")
		}
	}
	return tokenString, nil
}

func ParseDate(date string) (time.Time, error) {
	timeLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation("02-01-2006 15:04", date, timeLocation)
}

func DateToString(t time.Time) string {
	return t.Format("02-01-2006 15:04")
}
