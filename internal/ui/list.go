package ui

import (
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func newList(title string, termWidth, termHeight, headerHeight int) *widgets.List {
	l := widgets.NewList()
	l.Title = title
	l.TextStyle = termui.NewStyle(termui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, headerHeight, termWidth, termHeight-headerHeight)

	return l
}
