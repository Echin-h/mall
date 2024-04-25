package cache

import (
	"context"
	"fmt"
	"gin-mall/conf/sql"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitCache() {
	rConfig := sql.Config.Redis
	cl := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		//Username: rConfig.RedisUsername,
		//Password: rConfig.RedisPassword,
		DB: rConfig.RedisDbName,
	})
	_, err := cl.Ping(RedisContext).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	RedisClient = cl
}
