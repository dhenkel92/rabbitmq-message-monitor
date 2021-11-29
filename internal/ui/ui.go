package ui

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	uiRKList "github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/routing-key-list"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/search"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	headerHeight = 6
	searchHeigh  = 3
)

type UI struct {
	settingsHeader *widgets.Paragraph
	shortcutHeader *widgets.Paragraph
	list           *uiRKList.RoutingKeyList
	searchBar      *search.SearchBar

	previousKey    string
	isSearchActive bool
}

func NewUI() (*UI, error) {
	if err := termui.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize termui: %v", err)
	}

	termWidth, termHeight := termui.TerminalDimensions()

	settingsHeader := newSettingsHeader(types.WidgetSize{StartX: 0, StartY: 0, Width: termWidth / 100 * 30, Height: headerHeight})
	shortcutHeader := newShortcutHeader(types.WidgetSize{StartX: termWidth / 100 * 30, StartY: 0, Width: termWidth / 100 * 70, Height: headerHeight})
	searchBar := search.New(types.WidgetSize{StartX: 0, StartY: headerHeight, Width: termWidth, Height: (searchHeigh + headerHeight)})
	routingKeyList := uiRKList.New("Routing Keys", types.WidgetSize{StartX: 0, StartY: (headerHeight + searchHeigh), Width: termWidth, Height: termHeight})

	return &UI{
		settingsHeader: settingsHeader,
		shortcutHeader: shortcutHeader,
		list:           &routingKeyList,
		searchBar:      &searchBar,
	}, nil
}

func (ui *UI) Render() {
	termui.Render(ui.settingsHeader, ui.shortcutHeader, ui.searchBar.GetUiElement(), ui.list.GetUiElement())
}

func (ui *UI) UpdateRKListData(data []*collector.RoutingKeyStats) {
	ui.list.SetData(data)
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
