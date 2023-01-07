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

// [...] Login User
func (s *mainService) LoginUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var payload *model.UserSignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	var user model.User

	// Query username in DATABSER
	if err := db.Model(&user).Where("username = ?", payload.Username).Select(); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    "Invalid Username or Password",
			// "msg":    err.Error(),
		})
		return
	}

	// Query Password in DATABSER
	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    "Invalid Username or Password",
			// "msg":    err.Error(),
		})
		return
	}

	// config, _ := conf.LoadConfig(".")
	var tokenExpiresIn = viper.GetDuration(`TOKEN_EXPIRED_IN`)
	var tokenSecret = viper.GetString(`TOKEN_SECRET`)
	var tokenMaxAge = viper.GetInt(`TOKEN_MAXAGE`)

	// Generate Token
	token, err := utils.GenerateToken(tokenExpiresIn, user.ID, tokenSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
	}

	ctx.SetCookie("token", token, tokenMaxAge*60, "/", "localhost", false, true)

	// Old version
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(1440 * time.Minute).Unix(),
	// })
	// ss, err := token.SignedString([]byte("CPF-DataVisualization"))
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  token,
	})
}

// [...] SignOut User
func (s *mainService) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "success",
	})
}
