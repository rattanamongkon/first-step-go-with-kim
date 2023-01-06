package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDataPlantReq struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"is_active"`
	FactoryId int64  `json:"factory_id"`
}

type CreateDataPlantRes struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (s *mainService) CreatePlant(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	req := &CreateDataPlantReq{}
	res := &CreateDataPlantRes{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	plant := model.Plant{
		Name:      req.Name,
		Code:      req.Code,
		IsActive:  &req.IsActive,
		FactoryId: req.FactoryId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := plant.CreatePlant(db); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res.Status = true
	res.Msg = `success`
	ctx.JSON(http.StatusCreated, res)
}

func (s *mainService) ShowPlant(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var plants []model.Plant
	if err := db.Model(&plants).Select(); err != nil {
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
		"data": gin.H{
			"plants": plants,
		},
	})
}
