package redis

/*
 * Redis client for temp storage and cache.
 */
import (
	"log"

	"github.com/go-redis/redis"

	"echo-sample/conf"
)

var RedisClient *redis.Client

/*
 * Set connection to redis server
 */

func Init() {
	config := conf.LoadConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.GetString("redis.address"),
		DB:   config.GetInt("redis.db"),
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
