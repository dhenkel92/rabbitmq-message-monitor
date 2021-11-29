package uiRKList

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func New(title string, termWidth, termHeight, headerHeight int) RoutingKeyList {
	list := RoutingKeyList{
		list:    newList(title, termWidth, termHeight, headerHeight),
		data:    make([]*RoutingKeyData, 0),
		sorting: RK_DEFAULT_SORTING,
	}
	list.renderList()
	return list
}

func newList(title string, termWidth, termHeight, headerHeight int) *widgets.List {
	l := widgets.NewList()
	l.Title = title
	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, headerHeight, termWidth, termHeight-headerHeight)

	return l
}
