package rabbitmq

import (
	"github.com/streadway/amqp"
)

type ConsumeFunc func(msg amqp.Delivery)

func (consumer *Consumer) Consume(cb ConsumeFunc) error {
	msgs, err := consumer.channel.Consume(consumer.queue.Name, consumer.name, true, false, false, false, amqp.Table{})
	if err != nil {
		return err
	}

	for msg := range msgs {
		// todo: maybe add error handling
		cb(msg)
	}

	return nil
}
