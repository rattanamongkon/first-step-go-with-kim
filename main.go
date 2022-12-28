package main

import (
	"example/go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Register Service
	handler := controller.RegisterController()

	// Route
	route := gin.New()

	// Route API
	api := route.Group(`/api`)
	api.POST(`/user/create`, handler.CreateUser)

	//Run Gin
	route.Run()
}
