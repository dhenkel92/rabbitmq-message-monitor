package rabbitmq

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func NewConsumer(conf ConsumerConfig) (*Consumer, error) {
	consumer := Consumer{
		connection: nil,
		channel:    nil,
		queue:      nil,
		name:       conf.ConsumerName,
	}

	var err error

	consumer.connection, err = amqp.Dial(conf.ConnectionString)
	if err != nil {
		return nil, err
	}

	consumer.channel, err = consumer.connection.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := consumer.channel.QueueDeclare(
		fmt.Sprintf("%s-%s", conf.QueuePrefix, uuid.New().String()),
		true,
		true,
		false,
		false,
		amqp.Table{},
	)
	if err != nil {
		return nil, err
	}
	consumer.queue = &queue

	if err = consumer.createBindings(conf.Bindings); err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (consumer *Consumer) createBindings(bindings []settings.ExchangeRTBinding) error {
	for _, binding := range bindings {
		if err := consumer.channel.QueueBind(consumer.queue.Name, binding.RoutingKey, binding.ExchnageName, false, amqp.Table{}); err != nil {
			return err
		}
	}

	return nil
}
