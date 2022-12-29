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
	// var temp = model.User{}
	temp := model.User{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	// func (h *Handler) createBookHandler(c *gin.Context) {
	// 	var book Book

	// 	if err := c.ShouldBindJSON(&book); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	if err := ctx.ShouldBindJSON(&temp); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	// 	if result := h.db.Create(&book); result.Error != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": result.Error.Error(),
	// 		})
	// 		return
	// 	}

	// 	// books = append(books, book)
	// 	c.JSON(http.StatusCreated, book)
	// }

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
