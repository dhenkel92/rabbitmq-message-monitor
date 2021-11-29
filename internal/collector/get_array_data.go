package collector

func (collector *Collector) GetArrayData() []*RoutingKeyStats {
	vals := make([]*RoutingKeyStats, 0)
	// collector.mu.Lock()

	for _, entry := range collector.routingKeyStats {
		vals = append(vals, entry)
	}

	// collector.mu.Unlock()

	return vals
}
