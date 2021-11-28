package collector

import "sync"

type RoutingKeyStats struct {
	Count      int
	TotalBytes float64
}

// RoutingKeyStatsMap key is the routeingKey
type RoutingKeyStatsMap map[string]*RoutingKeyStats

type Collector struct {
	mu              sync.Mutex
	routingKeyStats RoutingKeyStatsMap
}
