package rabbitmq

import "testing"

func TestInitRabbitMQ(t *testing.T) {
	InitRabbitMQ()
	if GlobalRabbitMQ == nil {
		t.Error("GlobalRabbitMQ should not be nil")
	}
}
