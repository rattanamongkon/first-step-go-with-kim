package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDataMachineReq struct {
	CountFg   int64  `json:"count_fg"`
	CountNg   int64  `json:"count_ng"`
	TimeStamp string `json:"time_stamp"`
	FactoryId int64  `json:"factory_id"`
	PlantId   int64  `json:"plant_id"`
	StationId int64  `json:"station_id"`
	SkuId     int64  `json:"sku_id"`
}

type CreateDataMachineRes struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (s *mainService) CreateDataMachine(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	req := &CreateDataMachineReq{}
	res := &CreateDataMachineRes{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	machine := model.Machine{
		CountFg:   req.CountFg,
		CountNg:   req.CountNg,
		TimeStamp: req.TimeStamp,
		FactoryId: req.FactoryId,
		PlantId:   req.PlantId,
		StationId: req.StationId,
		SkuId:     req.SkuId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := machine.CreateMachine(db); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res.Status = true
	res.Msg = `success`
	ctx.JSON(http.StatusBadRequest, res)
}
