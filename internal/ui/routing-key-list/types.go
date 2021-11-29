package uiRKList

import (
	"github.com/gizak/termui/v3/widgets"
)

type RoutingKeySorting int

const (
	RK_DEFAULT_SORTING RoutingKeySorting = RK_SORTING_COUNT_DESC

	RK_SORTING_NAME_ASC        RoutingKeySorting = 0
	RK_SORTING_NAME_DESC       RoutingKeySorting = 1
	RK_SORTING_COUNT_ASC       RoutingKeySorting = 2
	RK_SORTING_COUNT_DESC      RoutingKeySorting = 3
	RK_SORTING_TOTAL_SIZE_ASC  RoutingKeySorting = 4
	RK_SORTING_TOTAL_SIZE_DESC RoutingKeySorting = 5
)

type RoutingKeyList struct {
	list    *widgets.List
	data    []*RoutingKeyData
	sorting RoutingKeySorting
}

func (list *RoutingKeyList) GetUiElement() *widgets.List {
	return list.list
}

type ListColumnWidhts struct {
	routingKey int
	count      int
	avg        int
	ninetyFife int
	ninetyNine int
	totalSize  int
}

type RoutingKeyData struct {
	RoutingKey string
	TotalBytes float64
	Count      int
	Avg        float64
	NinetyFife float64
	NinetyNine float64
}
