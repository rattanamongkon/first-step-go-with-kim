package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *mainService) CreateUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	// Test insert database
	temp := model.User{
		Name:      `name`,
		Email:     `ser`,
		Username:  `wer`,
		Password:  `sf`,
		IsActive:  false,
		Role:      `member`,
		FactoryId: 1,
		PlantId:   1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err := db.Model(&temp).Insert(); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"msg":    err.Error(),
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

}
