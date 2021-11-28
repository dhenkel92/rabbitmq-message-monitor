package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func (consumer *Consumer) Consume() error {
	msgs, err := consumer.channel.Consume(consumer.queue.Name, consumer.name, true, false, false, false, amqp.Table{})
	if err != nil {
		return err
	}

	for msg := range msgs {
		fmt.Println(msg.RoutingKey)
	}

	return nil
}
