package settings

type GenericSettings struct {
	ConnectionString string
	QueuePrefix      string
}

type ExchangeRTBinding struct {
	ExchnageName string
	RoutingKey   string
}

type ExchangeMonitoringSettings struct {
	Generic  GenericSettings
	Bindings []ExchangeRTBinding
}
