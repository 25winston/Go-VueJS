package redis

import (
	"log"
	"os"
	"strconv"
	"strings"
	"gmmshops.go/app/providers/logs"

	"github.com/go-redis/redis"
)

// Client :
type Client struct{ *redis.Client }

var c *Client

// Initial :
func Initial() {

	if strings.ToLower(os.Getenv("REDIS")) != "true" {
		log.Println("Redis: ", strings.ToLower(os.Getenv("REDIS")))
		log.Println("Redis: disabled")
		return
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		db = 0
	}

	c := Client{redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})}

	pong, err := c.Ping().Result()
	if err != nil {
		log.Fatal(err, pong)
		return
	}

	logs.Logger().Info("Redis: OK")

}

// GetClient :
func GetClient() *Client {
	if c == nil {
		Initial()
	}
	return c
}
