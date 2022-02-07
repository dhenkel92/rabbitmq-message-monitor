package cmd

import (
	"fmt"
	"time"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/history"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/rabbitmq"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
)

func Server(conf *settings.ServerSettings) error {
	consumer, err := rabbitmq.NewConsumer(rabbitmq.ConsumerConfig{
		ConnectionString: conf.Generic.ConnectionString,
		QueuePrefix:      conf.QueuePrefix,
		Bindings:         conf.Bindings,
		ConsumerName:     conf.Generic.ConsumerName,
	})
	if err != nil {
		return err
	}

	history := history.New()
	collector := collector.New()
	go consumer.Consume(createConsumeFunc(&collector))

	_ = helper.StartInterval(10*time.Second, updateHistory(&history, &collector))

	fmt.Println(conf)
	return nil
}
