package cmd

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/rabbitmq"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
)

func MonitorExchange(conf *settings.ExchangeMonitoringSettings) error {
	consumer, err := rabbitmq.NewConsumer(rabbitmq.ConsumerConfig{
		ConnectionString: conf.Generic.ConnectionString,
		QueuePrefix:      conf.QueuePrefix,
		Bindings:         conf.Bindings,
		ConsumerName:     conf.Generic.ConsumerName,
	})
	if err != nil {
		return err
	}

	if err = consumer.Consume(); err != nil {
		return err
	}

	return nil
}
