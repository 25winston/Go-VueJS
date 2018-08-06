package configs

import (
	"gmmshops.go/app"
	"gmmshops.go/app/routes"

	"github.com/iris-contrib/middleware/csrf"
)

// Configure is ...
func Configure(app *app.Application) {

	// app.Logger().Info("Application: initial")

	SetupLogs(app)

	// init gorm
	SetupDatabase(app)

	// setup session manager
	SetupSession(app)

	// setup Cache Service
	SetupCache(app)

	csrfProtect := csrf.Protect([]byte("9AB0F421E53A477C084477AEA06096F5"),
		csrf.Secure(false),
		csrf.MaxAge(3600),
		csrf.CookieName("X-CSRF-TOKEN"),
	)

	app.UseGlobal(csrfProtect)

	// define api route
	routes.SetupAPI(app)

	// define api route
	routes.SetupWEB(app)
}
