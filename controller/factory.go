package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDataFactoryReq struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	IsActive bool   `json:"is_active"`
}

type CreateDataFactoryRes struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	// Data   []model.Factory `json:"data"`
}

func (s *mainService) CreateFactory(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	req := &CreateDataFactoryReq{}
	res := &CreateDataFactoryRes{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	factory := model.Factory{
		Name:      req.Name,
		Code:      req.Code,
		IsActive:  req.IsActive,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := factory.CreateFactory(db); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"msg":    "success",
		"data":   factory,
	})
}

func (s *mainService) ShowFactory(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var factories []model.Factory

	if err := db.Model(&factories).Select(); err != nil {
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
		"data":   factories,
	})
}
