package settings

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func ParseExchangeSettingsFromCLI(c *cli.Context) (*ExchangeMonitoringSettings, error) {
	rawBindings := c.StringSlice("binding")
	bindings, err := parseBindings(rawBindings)
	if err != nil {
		return nil, err
	}

	return &ExchangeMonitoringSettings{
		Generic:  parseGenericSettingsFromCLI(c),
		Bindings: bindings,
	}, nil
}

func parseGenericSettingsFromCLI(c *cli.Context) GenericSettings {
	return GenericSettings{
		ConnectionString: c.String("connection-string"),
		QueuePrefix:      c.String("queue-prefix"),
	}
}

func parseBindings(raw []string) ([]ExchangeRTBinding, error) {
	result := make([]ExchangeRTBinding, len(raw))

	for i, entry := range raw {
		split := strings.Split(entry, "=")
		if len(split) != 2 {
			return result, fmt.Errorf("Invalid binding '%s'. Expected format <exchange>=<routing_key>", entry)
		}

		result[i] = ExchangeRTBinding{
			ExchnageName: split[0],
			RoutingKey:   split[1],
		}
	}

	return result, nil
}
