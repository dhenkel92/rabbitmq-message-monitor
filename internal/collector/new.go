package collector

func New() Collector {
	return Collector{
		routingKeyStats: make(RoutingKeyStatsMap),
		overallStats: OverallStats{
			TotalBytes:   0,
			MessageCount: 0,
		},
	}
}
