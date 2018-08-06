package controllers

import (
	"gmmshops.go/app"

	"github.com/kataras/iris/context"

	"github.com/kataras/iris/sessions"
)

// BaseController : base class Controller
type BaseController struct {
	Application *app.Application
}

// Session : Get Session
func (bc *BaseController) Session(ctx context.Context) *sessions.Session {
	return bc.Application.SessionManager.Start(ctx)
}
