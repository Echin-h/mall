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

// 不要觉得redis有多么高大上，可以把他想象成打怪升级的一个关卡
// 当用户访问的时候， 路由 - 中间件 - 业务逻辑 - redis - mysql
// 所以 redis 只是在 mysql 前面的一个东西而已，想什么数据放进去就放进去，随便放

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
