package sessions

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

var (
	// RedisDB : instance redis database
	RedisDB *redis.Database
)

// NewSessionRedis :
func NewSessionRedis() *sessions.Sessions {

	// RedisDB := RedisProvider.New()

	maxIdle, err := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	if err != nil {
		maxIdle = 0
	}

	maxActive, err := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	if err != nil {
		maxActive = 0
	}

	// replace with your running redis' server settings:
	RedisDB := redis.New(service.Config{
		Network:     "tcp",
		Addr:        os.Getenv("REDIS_ADDR"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		Database:    os.Getenv("REDIS_DATABASE"),
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(5) * time.Minute,
		Prefix:      os.Getenv("REDIS_PREFIX")}) // optionally configure the bridge between your redis server

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		if RedisDB != nil {
			RedisDB.Close()
		}
	})

	expires, err := strconv.ParseInt(os.Getenv("SESSION_EXPIRES"), 10, 64)
	if err != nil {
		expires = 60
	}

	var config sessions.Config

	if strings.ToLower(os.Getenv("SESSION_SECURE")) == "true" {
		hash := []byte(os.Getenv("SESSION_SECURE_HASH"))
		secret := []byte(os.Getenv("SESSION_SECURE_SECRET"))
		secure := securecookie.New(hash, secret)

		config = sessions.Config{
			Cookie:       os.Getenv("SESSION_NAME"),
			Expires:      time.Duration(expires) * time.Minute,
			Encode:       secure.Encode,
			Decode:       secure.Decode,
			AllowReclaim: true,
		}

	} else {
		config = sessions.Config{
			Cookie:       os.Getenv("SESSION_NAME"),
			Expires:      time.Duration(expires) * time.Minute,
			AllowReclaim: true,
		}
	}

	SessionManager := sessions.New(config)
	SessionManager.UseDatabase(RedisDB)

	return SessionManager
}
