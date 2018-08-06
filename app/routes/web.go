package routes

import (
	"gmmshops.go/app"

	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
)

// SetupWEB : config route
func SetupWEB(app *app.Application) {
	app.Use(logger.New())

	app.Favicon("./app/resources/favicon.ico")
	app.StaticWeb("/", "./public")

	app.RegisterView(iris.HTML("./app/views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./public/index.html", false)
	})
	app.Get("/{*}", func(ctx iris.Context) {
		ctx.ServeFile("./public/index.html", false)
	})
}
