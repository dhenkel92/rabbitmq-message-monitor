package collector

func (c *Collector) AppendRTStat(rt string, msgSize float64) {
	c.mu.Lock()

	if _, ok := c.routingKeyStats[rt]; !ok {
		c.routingKeyStats[rt] = &RoutingKeyStats{RoutingKey: rt, Count: 0, TotalBytes: 0}
	}

	c.routingKeyStats[rt].Count += 1
	c.routingKeyStats[rt].TotalBytes += msgSize

	c.mu.Unlock()
}
