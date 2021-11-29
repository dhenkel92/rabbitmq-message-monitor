package uiRKList

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/gizak/termui/v3/widgets"
)

type RoutingKeySorting int

const (
	RK_SORTING_NAME_ASC        RoutingKeySorting = 0
	RK_SORTING_NAME_DESC       RoutingKeySorting = 1
	RK_SORTING_COUNT_ASC       RoutingKeySorting = 2
	RK_SORTING_COUNT_DESC      RoutingKeySorting = 3
	RK_SORTING_TOTAL_SIZE_ASC  RoutingKeySorting = 4
	RK_SORTING_TOTAL_SIZE_DESC RoutingKeySorting = 5
)

type RoutingKeyList struct {
	list    *widgets.List
	data    []*collector.RoutingKeyStats
	sorting RoutingKeySorting
}

func (list *RoutingKeyList) GetUiElement() *widgets.List {
	return list.list
}
