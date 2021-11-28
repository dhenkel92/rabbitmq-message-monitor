package rabbitmq

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	"github.com/streadway/amqp"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	name       string
}

type ConsumerConfig struct {
	ConnectionString string
	QueuePrefix      string
	Bindings         []settings.ExchangeRTBinding
	ConsumerName     string
}
