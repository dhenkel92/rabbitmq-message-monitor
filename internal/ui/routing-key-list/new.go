package uiRKList

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func New(title string, size types.WidgetSize) RoutingKeyList {
	list := RoutingKeyList{
		list:    newList(title, size),
		data:    make([]*RoutingKeyData, 0),
		sorting: RK_DEFAULT_SORTING,
	}
	list.Render()
	return list
}

func newList(title string, size types.WidgetSize) *widgets.List {
	l := widgets.NewList()
	l.Title = title
	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	l.WrapText = false
	l.SetRect(size.StartX, size.StartY, size.EndX, size.EndY)

	return l
}
