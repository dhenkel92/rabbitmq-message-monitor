package collector

import "sync"

type RoutingKeyStats struct {
	RoutingKey string
	Count      int
	TotalBytes float64
	Sizes      []float64
}

// RoutingKeyStatsMap key is the routeingKey
type RoutingKeyStatsMap map[string]*RoutingKeyStats

type Collector struct {
	mu              sync.Mutex
	routingKeyStats RoutingKeyStatsMap
	overallStats    OverallStats
}

type OverallStats struct {
	TotalBytes   float64
	MessageCount int
}
