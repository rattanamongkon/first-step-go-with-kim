package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("CPF-DataVisualization"), nil
	})

	return err
}

func (s *mainService) AuthorizationMiddleware(c *gin.Context) {
	au := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(au, "Bearer ")

	if err := ValidateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
