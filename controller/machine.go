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

	// req := &CreateDataMachineReq{}
	// res := &CreateDataMachineRes{}

	var payload *model.MachineInput
	log.Println(payload)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	machine := model.Machine{
		CountFg:   payload.CountFg,
		CountNg:   payload.CountNg,
		TimeStamp: payload.TimeStamp,
		FactoryId: payload.FactoryId,
		PlantId:   payload.PlantId,
		StationId: payload.StationId,
		SkuId:     payload.SkuId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if _, err := db.Model(&machine).Insert(); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusConflict, gin.H{
			"status": false,
			"msg":    err.Error(),
		})
		return
	}

	// if err := machine.CreateMachine(db); err != nil {
	// 	log.Println(err.Error())
	// 	res.Msg = err.Error()
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// res.Status = true
	// res.Msg = `success`
	// ctx.JSON(http.StatusBadRequest, res)
	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		"msg":    "success",
	})
}

func (s *mainService) DataMachine(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db
	var machines []model.Machine

	if err := db.Model(&machines).Order("id ASC").Limit(50).Select(); err != nil {
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
			"machines": machines,
		},
	})
}
