package app

import (
	"time"
	"gmmshops.go/app/providers/logs"

	"github.com/kataras/golog"

	"github.com/jinzhu/gorm"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// Configurator : ...
type Configurator func(*Application)

// Application is ..
type Application struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	SessionManager *sessions.Sessions
	Gorm           *gorm.DB
	L              *golog.Logger
}

// NewApplication : returns a new Application.
func NewApplication(appName, appOwner string, configs ...Configurator) *Application {
	app := &Application{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, config := range configs {
		config(app)
	}
	logs.Logger().Info("Server: Running")
	return app
}

// Listen starts the http server with the specified "addr".
func (app *Application) Listen(addr string, cfgs ...iris.Configurator) {
	app.Run(iris.Addr(addr), cfgs...)
}
