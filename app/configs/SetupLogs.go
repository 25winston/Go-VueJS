package configs

import (
	"os"
	"gmmshops.go/app"
	"gmmshops.go/app/providers/logs"
)

// SetupLogs : setup logger
func SetupLogs(app *app.Application) {
	var logger = logs.Logger()
	app.Logger().Install(logger)
	app.Logger().SetLevel(os.Getenv("LOG_LEVEL"))
	app.Logger().Info("Application Logger: Active")
}
