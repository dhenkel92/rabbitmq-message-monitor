package ui

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/collector"
	uiRKList "github.com/dhenkel92/rabbitmq-message-monitor/internal/ui/routing-key-list"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const headerHeight = 6

type UI struct {
	settingsHeader *widgets.Paragraph
	shortcutHeader *widgets.Paragraph
	list           *uiRKList.RoutingKeyList

	previousKey string
}

func NewUI() (*UI, error) {
	if err := termui.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize termui: %v", err)
	}

	termWidth, termHeight := termui.TerminalDimensions()

	settingsHeader := newSettingsHeader(WidgetSize{startX: 0, startY: 0, width: termWidth / 100 * 30, height: headerHeight})
	shortcutHeader := newShortcutHeader(WidgetSize{startX: termWidth / 100 * 30, startY: 0, width: termWidth / 100 * 70, height: headerHeight})
	routingKeyList := uiRKList.New("Routing Keys", termWidth, termHeight, headerHeight)

	return &UI{
		settingsHeader: settingsHeader,
		shortcutHeader: shortcutHeader,
		list:           &routingKeyList,
	}, nil
}

func (ui *UI) Render() {
	termui.Render(ui.settingsHeader, ui.shortcutHeader, ui.list.GetUiElement())
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
