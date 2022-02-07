package settings

import "github.com/urfave/cli/v2"

type ServerSettings struct {
	Generic     GenericSettings
	Bindings    []ExchangeRTBinding
	QueuePrefix string
}

func ParseServerSettings(cli *cli.Context) (*ServerSettings, error) {
	rawBindings := cli.StringSlice("binding")
	bindings, err := parseBindings(rawBindings)
	if err != nil {
		return nil, err
	}

	return &ServerSettings{
		Generic:     parseGenericSettingsFromCLI(cli),
		Bindings:    bindings,
		QueuePrefix: cli.String("queue-prefix"),
	}, nil
}
