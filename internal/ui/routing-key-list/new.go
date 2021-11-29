package uiRKList

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func New(title string, termWidth, termHeight, headerHeight int) RoutingKeyList {
	return RoutingKeyList{
		list:    newList(title, termWidth, termHeight, headerHeight),
		data:    make([]*collector.RoutingKeyStats, 0),
		sorting: RK_SORTING_COUNT_DESC,
	}
}

func newList(title string, termWidth, termHeight, headerHeight int) *widgets.List {
	l := widgets.NewList()
	l.Title = title
	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, headerHeight, termWidth, termHeight-headerHeight)

	return l
}
