package caches

// CacheService : cache service interface
type CacheService interface {
	Get(key string) (string, error)
	Set(key string, value string, expires int64) error
	Delete(key string) error
	Flush(key string) error
	Expire(key string, expires int64) error
	Exists(key string) bool
}
