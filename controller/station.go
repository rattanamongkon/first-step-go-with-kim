package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *mainService) CreateStation(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	// Test insert database
	tmp := model.Station{
		ID:        1,
		Name:      `name`,
		Code:      `LLK`,
		IsActive:  false,
		FactoryId: 1,
		PlantId:   1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err := db.Model(&tmp).Insert(); err != nil {
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