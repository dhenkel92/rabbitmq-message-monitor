package uiRKList

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/montanaflynn/stats"
)

func (list *RoutingKeyList) SetData(data []*collector.RoutingKeyStats) {
	converted := make([]*RoutingKeyData, 0)
	for _, rkStats := range data {
		converted = append(converted, rkDataFromRKStats(rkStats))
	}

	list.data = sortData(list.sorting, converted)
	list.renderList()
}

func rkDataFromRKStats(rkStats *collector.RoutingKeyStats) *RoutingKeyData {
	avg, _ := stats.Mean(rkStats.Sizes)
	ninetyFife, _ := stats.Percentile(rkStats.Sizes, 95)
	ninetyNine, _ := stats.Percentile(rkStats.Sizes, 99)
	return &RoutingKeyData{
		RoutingKey: rkStats.RoutingKey,
		Count:      rkStats.Count,
		TotalBytes: rkStats.TotalBytes,
		Avg:        avg,
		NinetyFife: ninetyFife,
		NinetyNine: ninetyNine,
	}
}
