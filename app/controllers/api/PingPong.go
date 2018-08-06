package api

import (
	"gmmshops.go/app"
	"gmmshops.go/app/controllers"
	"gmmshops.go/app/models"

	"github.com/kataras/iris/context"
)

// PingPong : Sample Controller Class
type PingPong struct {
	controllers.BaseController
}

// NewPinoPong : create new instance PingPong  Controller Class
func NewPinoPong(app *app.Application) *PingPong {
	return &PingPong{controllers.BaseController{Application: app}}
}

// Ping : Sample Controller Method
func (pp *PingPong) Ping(ctx context.Context) {
	Session := pp.Session(ctx)
	Session.Set("MSG", "PONG")

	ctx.JSON(models.NewResponse("OK", "Pong", nil))
}
