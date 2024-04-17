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
