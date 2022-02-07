package settings

import "github.com/urfave/cli/v2"

type ExchangeMonitoringSettings struct {
	Generic     GenericSettings
	Bindings    []ExchangeRTBinding
	QueuePrefix string
}

func ParseExchangeSettingsFromCLI(c *cli.Context) (*ExchangeMonitoringSettings, error) {
	rawBindings := c.StringSlice("binding")
	bindings, err := parseBindings(rawBindings)
	if err != nil {
		return nil, err
	}

	return &ExchangeMonitoringSettings{
		Generic:     parseGenericSettingsFromCLI(c),
		Bindings:    bindings,
		QueuePrefix: c.String("queue-prefix"),
	}, nil
}
