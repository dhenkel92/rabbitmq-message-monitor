package main

import (
	"log"
	"os"

	"github.com/dhenkel92/rabbitmq-message-monitor/cmd"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/settings"
	cli "github.com/urfave/cli/v2"
)

func initApp() *cli.App {
	return &cli.App{
		Name:  "RabbitMQ Message Monitor",
		Usage: "Different ways of how to monitor messages that are going through your RabbitMQ infrastructure.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "connection-string",
				Usage:   "",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "consumer-name",
				Aliases: []string{"name"},
				Value:   "rmq-message-monitor",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "monitor-exchange",
				Usage:   "Bind to several exchange and visualize the different messages.",
				Aliases: []string{"exchange"},
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     "binding",
						Aliases:  []string{"b"},
						Usage:    "Which bindings should be created for monitoring. Format: <exchange>=<routing_key>",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "queue-prefix",
						Usage: "Name prefix to use for temporal queues.",
						Value: "rmq-message-monitor",
					},
				},
				Action: func(c *cli.Context) error {
					conf, err := settings.ParseExchangeSettingsFromCLI(c)
					if err != nil {
						return err
					}

					return cmd.MonitorExchange(conf)
				},
			},
			{
				Name:  "server",
				Usage: "Binds to several exchanges and exports the data to /metrics",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     "binding",
						Aliases:  []string{"b"},
						Usage:    "Which bindings should be created for monitoring. Format: <exchange>=<routing_key>",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					conf, err := settings.ParseServerSettings(c)
					if err != nil {
						return err
					}

					return cmd.Server(conf)
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
