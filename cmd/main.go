package main

import (
	"fmt"
	"gin-mall/conf/sql"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/cache"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/rabbitmq"
	"gin-mall/router"
)

// @title mall
// @version 1.0
// @description gin mall for learning
// @termsOfService http://swagger.io/terms/

// @contact.name echin
// @contact.url https://echin-h.github.io/
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {
	loading()
	r := router.NewRouter()
	_ = r.Run(sql.Config.System.HttpPort)
	fmt.Println("启动成功...")
}

func loading() {
	sql.InitConfig()
	dao.InitMysql()
	cache.InitCache()
	rabbitmq.InitRabbitMQ()
	log.InitLog()
}
