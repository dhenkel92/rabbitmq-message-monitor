package cmd

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/rabbitmq"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	"github.com/streadway/amqp"
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

	collector := collector.New()
	if err = consumer.Consume(createConsumeFunc(&collector)); err != nil {
		return err
	}

	return nil
}

func createConsumeFunc(collect *collector.Collector) rabbitmq.ConsumeFunc {
	return func(msg amqp.Delivery) {
		collect.AppendRTStat(msg.RoutingKey, float64(len(msg.Body)))
		fmt.Println(msg.RoutingKey)
	}
}
