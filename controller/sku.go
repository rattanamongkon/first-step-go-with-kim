package controller

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDataSkuReq struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateDataSkuRes struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (s *mainService) CreateSku(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	req := &CreateDataSkuReq{}
	res := &CreateDataSkuRes{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	sku := model.Sku{
		Name:      req.Name,
		Code:      req.Code,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := sku.CreateSku(db); err != nil {
		log.Println(err.Error())
		res.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res.Status = true
	res.Msg = `success`
	ctx.JSON(http.StatusCreated, res)
}

func (s *mainService) ShowSku(ctx *gin.Context) {
	db := connection.GetConnectionDB()
	_ = db

	var sku []model.Sku

	if err := db.Model(&sku).Select(); err != nil {
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
			"sku": sku,
		},
	})
}
