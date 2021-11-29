package uiRKList

import (
	"github.com/gizak/termui/v3/widgets"
)

type RoutingKeyList struct {
	list *widgets.List
	data []*RoutingKeyData

	sorting              RoutingKeySorting
	nameFilterExpression string
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
