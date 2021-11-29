package cmd

import (
	"time"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/rabbitmq"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/ui"
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

	ui, err := ui.NewUI()
	collector := collector.New()
	// todo: get error
	go consumer.Consume(createConsumeFunc(&collector))
	helper.StartInterval(10*time.Second, updateUiFnFactory(ui, &collector))

	if err != nil {
		return err
	}
	if err = ui.Start(); err != nil {
		return err
	}

	return nil
}

func updateUiFnFactory(ui *ui.UI, collector *collector.Collector) helper.IntervalAction {
	return func() {
		ui.UpdateListEntries(collector.ToStrRows(ui.LineLength()))
	}
}

func createConsumeFunc(collect *collector.Collector) rabbitmq.ConsumeFunc {
	return func(msg amqp.Delivery) {
		collect.AppendRTStat(msg.RoutingKey, float64(len(msg.Body)))
	}
}
