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
	route.POST(`/login`, handler.LoginUser)
	route.GET(`/logout`, handler.LogoutUser)

	api := route.Group(`/api`, handler.DeserializeUser)
	// factory
	api.GET(`/factory`, handler.ShowFactory)
	api.POST(`/factory/create`, handler.CreateFactory)

	// plant
	api.GET(`/plant`, handler.ShowPlant)
	api.POST(`/plant/create`, handler.CreatePlant)

	// station
	api.GET(`/station`, handler.ShowStation)
	api.POST(`/station/create`, handler.CreateStation)

	// sku
	api.GET(`/sku`, handler.ShowSku)
	api.POST(`/sku/create`, handler.CreateSku)

	// user
	api.GET(`/user`, handler.ShowUser)
	api.POST(`/user/create`, handler.CreateUser)

	// machine
	route.POST(`/machine`, handler.CreateDataMachine)
	api.GET(`/machine`, handler.DataMachine)

	//Run Gin
	route.Run()
}
