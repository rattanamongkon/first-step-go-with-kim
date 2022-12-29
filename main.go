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
	api.GET(`/user`, handler.ShowUser)
	api.POST(`/user/create`, handler.CreateUser)
	api.POST(`/machine/create`, handler.CreateDataMachine)

	//Run Gin
	route.Run()
}
