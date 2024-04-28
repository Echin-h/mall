package rabbitmq

import (
	conf "gin-mall/conf/sql"
	amqp "github.com/rabbitmq/amqp091-go"
	"strings"
)

var GlobalRabbitMQ *amqp.Connection

// InitRabbitMQ initializes the rabbitMQ connection in the middleware
func InitRabbitMQ() {
	rConfig := conf.Config.RabbitMq
	pathRabbitMQ := strings.Join([]string{rConfig.RabbitMQ, "://", rConfig.RabbitMQUser, ":",
		rConfig.RabbitMQPassWord, "@", rConfig.RabbitMQHost, ":", rConfig.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ)
	if err != nil {
		panic(err)
	}
	GlobalRabbitMQ = conn
}
