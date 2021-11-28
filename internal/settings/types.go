package settings

type GenericSettings struct {
	ConnectionString string
	ConsumerName     string
}

type ExchangeRTBinding struct {
	ExchnageName string
	RoutingKey   string
}

type ExchangeMonitoringSettings struct {
	Generic     GenericSettings
	Bindings    []ExchangeRTBinding
	QueuePrefix string
}
