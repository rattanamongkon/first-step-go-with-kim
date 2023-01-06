package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"example/go-rest-api/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func (s *mainService) DeserializeUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db
	var token string
	cookie, err := ctx.Cookie("token")

	authorizationHeader := ctx.Request.Header.Get("Authorization")

	fields := strings.Fields(authorizationHeader)

	// If token is missing
	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"msg":    "Authentication Failed, Token is missing",
		})
		return
	}

	if len(fields) != 0 && fields[0] == "Bearer" {
		token = fields[1]
	} else if err == nil {
		token = cookie
	}

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"msg":    "You are not logged in",
		})
		return
	}

	var tokenSecret = viper.GetString(`TOKEN_SECRET`)
	sub, err := utils.ValidateToken(token, tokenSecret)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	var user model.User

	if err := db.Model(&user).Where("id = ?", fmt.Sprint(sub)).Select(); err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status": false,
			"msg":    "The Username belonging to this token no logger exists",
		})
		return
	}
	ctx.Set("currentUser", user)
	ctx.Next()
}
