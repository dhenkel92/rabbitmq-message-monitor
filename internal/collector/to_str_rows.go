package collector

import (
	"fmt"
	"sort"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
)

func (collector *Collector) ToStrRows(rowLenght int) []string {
	vals := make([]RoutingKeyStats, 0)
	collector.mu.Lock()
	for _, entry := range collector.routingKeyStats {
		vals = append(vals, *entry)
	}
	collector.mu.Unlock()

	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].Count > vals[j].Count
	})

	rtL, countL, sizeL := calculateCellWidths(rowLenght)

	result := make([]string, 0)
	result = append(result, fmt.Sprintf("[%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white)", rtL, "Routing Key", countL, "Count", sizeL, "Total Size"))
	for _, entry := range vals {
		result = append(result, fmt.Sprintf("%-*s %-*d %-*s", rtL, entry.RoutingKey, countL, entry.Count, sizeL, helper.FormatBytesToMb(entry.TotalBytes)))
	}

	return result
}

func calculateCellWidths(rowLenght int) (int, int, int) {
	routingKey := rowLenght / 100 * 60
	count := rowLenght / 100 * 20
	totalSize := rowLenght / 100 * 20
	return routingKey, count, totalSize
}
