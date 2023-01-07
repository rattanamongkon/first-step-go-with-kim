package main

import (
	"example/go-rest-api/conf"
	"example/go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//init func
	conf.Init()

	// Route
	app := gin.New()
	routes.Router(app)

	//Run Gin
	app.Run()
}
