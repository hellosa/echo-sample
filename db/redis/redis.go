package redis

/*
 * Redis client for temp storage and cache.
 */
import (
	"log"

	"github.com/go-redis/redis"

	"data-service/conf"
)

var RedisClient *redis.Client

/*
 * Set connection to redis server
 */

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: conf.ADDRESS,
		DB:   conf.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
