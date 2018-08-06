package redis

import (
	"time"
	"gmmshops.go/app/providers/redis"
)

// CacheService :
type CacheService struct {
	client *redis.Client
}

// Get :
func (c *CacheService) Get(key string) (string, error) {
	return c.client.Get("key").Result()
}

// Set :
func (c *CacheService) Set(key string, value string, expires int64) error {
	return c.client.Set(key, value, time.Duration(expires)*time.Minute).Err()
}

// Delete :
func (c *CacheService) Delete(key string) error {
	return c.client.Del(key).Err()
}

// Flush :
func (c *CacheService) Flush(key string) error {
	return c.client.ExpireAt(key, time.Now()).Err()
}

// Expire :
func (c *CacheService) Expire(key string, expires int64) error {
	return c.client.Expire(key, time.Duration(expires)*time.Minute).Err()
}

// Exists :
func (c *CacheService) Exists(key string) bool {
	_, err := c.Get(key)
	if err != nil {
		return false
	}
	return true
}

// New :
func New() CacheService {
	return CacheService{client: redis.GetClient()}
}
