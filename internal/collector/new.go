package collector

func New() Collector {
	return Collector{
		routingKeyStats: make(RoutingKeyStatsMap),
	}
}
