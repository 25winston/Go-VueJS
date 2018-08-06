package configs

import (
	"gmmshops.go/app"

	"github.com/kataras/iris"
)

// SetupViews :
func SetupViews(app *app.Application) {
	app.RegisterView(iris.HTML("/app/views", ".html").Layout("layouts/layout.html"))

}
