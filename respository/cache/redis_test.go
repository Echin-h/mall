package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
)

func TestInitCache(t *testing.T) {
	cl := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
		DB:   4,
	})
	_, err := cl.Ping(RedisContext).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
