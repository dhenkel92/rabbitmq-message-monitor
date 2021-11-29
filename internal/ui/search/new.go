package search

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	"github.com/gizak/termui/v3/widgets"
)

func New(size types.WidgetSize) SearchBar {
	bar := SearchBar{
		uiElement:  newSettingsHeader(size),
		searchTerm: "",
	}
	bar.Render()
	return bar
}

func newSettingsHeader(size types.WidgetSize) *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Title = "Search"
	header.Text = "hello search"
	header.SetRect(size.StartX, size.StartY, size.Width, size.Height)
	header.Border = true
	return header
}
