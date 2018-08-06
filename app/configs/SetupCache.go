package configs

import (
	"gmmshops.go/app"
	"gmmshops.go/app/providers/services"
	"gmmshops.go/app/providers/services/caches/redis"
)

// SetupCache : Cache Provider
func SetupCache(app *app.Application) {
	app.Logger().Info("Service: Loading")
	services.Register("caches.CacheService", redis.New())
	app.Logger().Info("Service: Completed")
}
