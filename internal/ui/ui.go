package ui

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	lineChart "github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/line-chart"
	uiRKList "github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/routing-key-list"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/search"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	settingsHeight = 6
	shortcutHeight = 8
	searchHeigh    = 3
)

type UI struct {
	settingsHeader *widgets.Paragraph
	shortcutHeader *widgets.Paragraph
	list           *uiRKList.RoutingKeyList
	searchBar      *search.SearchBar

	msgSizeHistory  *lineChart.LineChart
	msgCountHistory *lineChart.LineChart

	previousKey    string
	isSearchActive bool
}

func NewUI() (*UI, error) {
	if err := termui.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize termui: %v", err)
	}

	termWidth, termHeight := termui.TerminalDimensions()
	settingsRightBorder := int(float64(termWidth) / 100.0 * 30.0)
	msgSizeHistRightBorder := int(float64(termWidth) / 100.0 * 65.0)

	settingsHeader := newSettingsHeader(types.WidgetSize{StartX: 0, StartY: 0, EndX: settingsRightBorder, EndY: settingsHeight})
	shortcutHeader := newShortcutHeader(types.WidgetSize{StartX: 0, StartY: settingsHeight, EndX: settingsRightBorder, EndY: (settingsHeight + shortcutHeight)})

	msgSizeHistory := lineChart.New(
		"Total Message Size (MB / 10s)",
		types.WidgetSize{StartX: settingsRightBorder, EndX: msgSizeHistRightBorder, StartY: 0, EndY: (settingsHeight + shortcutHeight)},
	)
	msgCountHistory := lineChart.New(
		"Total Message Count (k / s)",
		types.WidgetSize{StartX: msgSizeHistRightBorder, EndX: termWidth, StartY: 0, EndY: (settingsHeight + shortcutHeight)},
	)

	searchBar := search.New(types.WidgetSize{StartX: 0, StartY: (settingsHeight + shortcutHeight), EndX: termWidth, EndY: (searchHeigh + settingsHeight + shortcutHeight)})
	routingKeyList := uiRKList.New("Routing Keys", types.WidgetSize{StartX: 0, StartY: (settingsHeight + searchHeigh + shortcutHeight), EndX: termWidth, EndY: termHeight})

	return &UI{
		settingsHeader:  settingsHeader,
		shortcutHeader:  shortcutHeader,
		list:            &routingKeyList,
		searchBar:       &searchBar,
		msgSizeHistory:  &msgSizeHistory,
		msgCountHistory: &msgCountHistory,
	}, nil
}

func (ui *UI) Render() {
	termui.Render(ui.settingsHeader, ui.shortcutHeader, ui.msgSizeHistory.GetUiElement(), ui.msgCountHistory.GetUiElement(), ui.searchBar.GetUiElement(), ui.list.GetUiElement())
}

func (ui *UI) UpdateData(data []*collector.RoutingKeyStats, msgSizeHistory, msgCount []float64) {
	ui.list.SetData(data)
	ui.msgSizeHistory.SetBytes(msgSizeHistory)
	ui.msgCountHistory.SetRPS(msgCount)
	ui.Render()
}

func (ui *UI) Start() error {
	defer termui.Close()

	ui.Render()
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		if ui.handleKeyPress(&e) {
			return nil
		}
	}
}
