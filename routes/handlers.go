package routes

import (
	"example/go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	handler := controller.RegisterController()

	// Route API
	app.POST(`/login`, handler.LoginUser)
	app.GET(`/logout`, handler.LogoutUser)

	api := app.Group(`/api`, handler.DeserializeUser)
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
	app.POST(`/machine`, handler.CreateDataMachine)
	api.GET(`/machine`, handler.DataMachine)

}
