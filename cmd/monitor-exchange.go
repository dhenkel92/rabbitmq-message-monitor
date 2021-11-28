package cmd

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
)

func MonitorExchange(conf *settings.ExchangeMonitoringSettings) error {
	fmt.Println(*conf)
	return nil
}
