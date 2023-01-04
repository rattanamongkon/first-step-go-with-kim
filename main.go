package main

import (
	"example/go-rest-api/conf"
	"example/go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	//init func
	conf.Init()

	// Register Service
	handler := controller.RegisterController()

	// Route
	route := gin.New()

	// Route API
	api := route.Group(`/api`)
	// factory
	api.GET(`/factory`, handler.ShowFactory)
	api.POST(`/factory/create`, handler.CreateFactory)

	// plant
	api.POST(`/plant`, handler.ShowPlant)
	api.POST(`/plant/create`, handler.CreatePlant)

	// station
	api.POST(`/station`, handler.ShowStation)
	api.POST(`/station/create`, handler.CreateStation)

	// sku
	api.POST(`/sku`, handler.ShowSku)
	api.POST(`/sku/create`, handler.CreateSku)

	// user
	api.GET(`/user`, handler.ShowUser)
	api.POST(`/user/create`, handler.CreateUser)

	// machine
	api.POST(`/machine/create`, handler.CreateDataMachine)

	//Run Gin
	route.Run()
}
