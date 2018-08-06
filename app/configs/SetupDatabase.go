package configs

import (
	"os"
	"gmmshops.go/app"
	"gmmshops.go/app/providers/orm/gorm"
)

// SetupDatabase :
func SetupDatabase(app *app.Application) {
	if (os.Getenv("DATABASE") == "mysql" || os.Getenv("DATABASE") == "mssql" || os.Getenv("DATABASE") == "postgres") && os.Getenv("DATABASE_CONN") != "" {
		gorm.Gorm()
		app.Logger().Info("Database: OK")
	} else {
		app.Logger().Info("Database: Disable")
	}

}
