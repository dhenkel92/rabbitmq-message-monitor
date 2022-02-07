package settings

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

type GenericSettings struct {
	ConnectionString string
	ConsumerName     string
}

type ExchangeRTBinding struct {
	ExchnageName string
	RoutingKey   string
}

func parseGenericSettingsFromCLI(c *cli.Context) GenericSettings {
	return GenericSettings{
		ConnectionString: c.String("connection-string"),
		ConsumerName:     c.String("consumer-name"),
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
