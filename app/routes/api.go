package routes

import (
	"gmmshops.go/app"
	"gmmshops.go/app/controllers/api"
)

// SetupAPI : config route
func SetupAPI(app *app.Application) {

	// routing for `/api`
	RouteAPI := app.Party("/api")

	// routing for `/api/v1`
	APIV1 := RouteAPI.Party("/v1")

	PingPongController := api.NewPinoPong(app)
	APIV1.Get("/ping", PingPongController.Ping)

}
