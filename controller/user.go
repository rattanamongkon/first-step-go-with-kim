package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"

	"github.com/gin-gonic/gin"
)

func (s *mainService) CreateUser(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	temp := model.User{
		ID:        1,
		Name:      `name`,
		Email:     `ser`,
		Username:  `wer`,
		Password:  `sf`,
		IsActive:  false,
		Role:      `member`,
		FactoryId: 1,
		PlantId:   1,
	}

	if _, err := db.Model(&temp).Insert(); err != nil {
		log.Println(err.Error())
		ctx.JSON(500, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"msg":    "success",
	})
}
