package rabbitmq

//func SendMessage(ctx context.Context, rabbitMqQueue string, message []byte) error {
//	ch, err := GlobalRabbitMQ.Channel()
//	if err != nil {
//		return err
//	}
//	queue, err := ch.QueueDeclare(rabbitMqQueue, false, false, false, false, nil)
//	if err != nil {
//		return err
//	}
//	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: message})
//	return err
//}
//
//func ConsumeMessage(ctx context.Context, rabbitMqQueue string) (<-chan amqp.Delivery, error) {
//	ch, err := GlobalRabbitMQ.Channel()
//	if err != nil {
//		return nil, err
//	}
//	q, err := ch.QueueDeclare(rabbitMqQueue, false, false, false, false, nil)
//	if err != nil {
//		return nil, err
//	}
//	return ch.Consume(q.Name, "", true, false, false, false, nil)
//}
