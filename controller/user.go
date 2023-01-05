package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"example/go-rest-api/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (s *mainService) CreateUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var payload *model.UserSingUpInput
	// Check json fomat
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	user := model.User{
		Name:      payload.Name,
		Username:  payload.Username,
		Password:  hashedPassword,
		Email:     payload.Email,
		Role:      payload.Role,
		IsActive:  &payload.IsActive,
		FactoryId: payload.FactoryId,
		PlantId:   payload.PlantId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	// Insert user to DATABSER
	if _, err := db.Model(&user).Insert(); err != nil {
		log.Println("Insert user to DATABSER:" + err.Error())
		ctx.JSON(http.StatusConflict, gin.H{
			"status": false,
			"msg":    "User with that email already exists",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"msg":    "success",
	})
}

func (s *mainService) ShowUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var temp []model.User

	if err := db.Model(&temp).Select(); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "success",
		"data":   temp,
	})
}

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
			"msg":    "Invalid email or Password",
			// "msg":    err.Error(),
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1440 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte("CPF-DataVisualization"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  ss,
	})
}
