package main

import (
	"log"
	"os"

	"github.com/dhenkel92/rabbitmq-message-monitor/cmd"
	cli "github.com/urfave/cli/v2"
)

func initApp() *cli.App {
	return &cli.App{
		Name:  "RabbitMQ Message Monitor",
		Usage: "",
		Commands: []*cli.Command{
			{
				Name:    "monitor-exchange",
				Aliases: []string{"exchange"},
				Action: func(c *cli.Context) error {
					return cmd.MonitorExchange()
				},
			},
		},
	}
}

func main() {
	app := initApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
