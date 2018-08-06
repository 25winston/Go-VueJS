package configs

import (
	"os"
	"strings"
	"gmmshops.go/app"

	config "gmmshops.go/app/configs/sessions"

	"github.com/kataras/iris/sessions"
)

var (
	// SessionManager : instance iris sessions
	SessionManager *sessions.Sessions
)

// SetupSession :
func SetupSession(app *app.Application) {
	if strings.ToLower(os.Getenv("SESSION")) == "redis" {
		app.SessionManager = config.NewSessionRedis()

	} else {
		// default is cookies
		app.SessionManager = config.NewSessionCookie()
	}
	app.Logger().Info("Session: OK")
}
