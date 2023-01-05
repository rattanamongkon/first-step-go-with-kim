package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDataStationReq struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"is_active"`
	PlantId   int64  `json:"plant_id"`
	FactoryId int64  `json:"factory_id"`
}

type CreateDataStationRes struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (s *mainService) CreateStation(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db
	req := &CreateDataStationReq{}
	res := &CreateDataStationRes{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	station := model.Station{
		Name:      req.Name,
		Code:      req.Code,
		IsActive:  req.IsActive,
		PlantId:   req.PlantId,
		FactoryId: req.FactoryId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := station.CreateStation(db); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res.Status = true
	res.Msg = `success`
	ctx.JSON(http.StatusCreated, res)
}

func (s *mainService) ShowStation(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var station []model.Station

	if err := db.Model(&station).Select(); err != nil {
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
		"data":   station,
	})
}
