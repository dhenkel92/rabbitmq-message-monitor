package cmd

import (
	"time"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/rabbitmq"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	"github.com/gosuri/uilive"
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
	// todo: get error
	go consumer.Consume(createConsumeFunc(&collector))

	writer := uilive.New()
	writer.Start()
	helper.ClearTerminal()
	for {
		collector.PrintTable(writer)
		time.Sleep(10 * time.Second)
	}
}

func createConsumeFunc(collect *collector.Collector) rabbitmq.ConsumeFunc {
	return func(msg amqp.Delivery) {
		collect.AppendRTStat(msg.RoutingKey, float64(len(msg.Body)))
	}
}
