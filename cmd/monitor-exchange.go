package cmd

import (
	"time"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/history"
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
	history := history.New()

	// todo: get error
	go consumer.Consume(createConsumeFunc(&collector))
	stopUpdateUI := helper.StartInterval(10*time.Second, updateUiFnFactory(ui, &collector, &history))
	stopUpdateHistory := helper.StartInterval(10*time.Second, updateHistory(&history, &collector))

	if err != nil {
		return err
	}
	if err = ui.Start(); err != nil {
		return err
	}

	stopUpdateUI <- true
	stopUpdateHistory <- true

	return nil
}

func updateUiFnFactory(ui *ui.UI, collector *collector.Collector, recorder *history.HistoryRecorder) helper.IntervalAction {
	return func() {
		ui.UpdateData(
			collector.GetArrayData(),
			recorder.GetRecordValues(history.RECORD_SUM_SIZE_KEY),
			recorder.GetRecordValues(history.RECORD_MSG_COUNT_KEY),
		)
	}
}

func updateHistory(recorder *history.HistoryRecorder, collector *collector.Collector) helper.IntervalAction {
	return func() {
		overallStats := collector.GetOverallStats()
		recorder.RecordDiff(history.RECORD_SUM_SIZE_KEY, overallStats.TotalBytes)
		recorder.RecordDiff(history.RECORD_MSG_COUNT_KEY, float64(overallStats.MessageCount))
	}
}

func createConsumeFunc(collect *collector.Collector) rabbitmq.ConsumeFunc {
	return func(msg amqp.Delivery) {
		collect.AppendRTStat(msg.RoutingKey, float64(len(msg.Body)))
	}
}
