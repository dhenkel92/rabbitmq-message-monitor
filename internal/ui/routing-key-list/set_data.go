package uiRKList

import "github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"

func (list *RoutingKeyList) SetData(data []*collector.RoutingKeyStats) {
	list.data = sortData(list.sorting, data)
	list.renderList()
}
